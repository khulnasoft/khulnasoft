use crate::native::db::connection::KhulnasoftDbConnection;
use rusqlite::{Connection, OpenFlags};
use std::fs::{remove_file, File};
use std::path::{Path, PathBuf};
use tracing::{debug, trace};

pub(super) struct LockFile {
    file: File,
    path: PathBuf,
}

pub(super) fn unlock_file(lock_file: &LockFile) {
    if lock_file.path.exists() {
        fs4::fs_std::FileExt::unlock(&lock_file.file)
            .and_then(|_| remove_file(&lock_file.path))
            .ok();
    }
}

pub(super) fn create_lock_file(db_path: &Path) -> anyhow::Result<LockFile> {
    let lock_file_path = db_path.with_extension("lock");
    let lock_file = File::create(&lock_file_path)
        .map_err(|e| anyhow::anyhow!("Unable to create db lock file: {:?}", e))?;

    trace!("Getting lock on db lock file");
    fs4::fs_std::FileExt::lock_exclusive(&lock_file)
        .inspect(|_| trace!("Got lock on db lock file"))
        .map_err(|e| anyhow::anyhow!("Unable to lock the db lock file: {:?}", e))?;
    Ok(LockFile {
        file: lock_file,
        path: lock_file_path,
    })
}

pub(super) fn initialize_db(
    khulnasoft_version: String,
    db_path: &Path,
) -> anyhow::Result<KhulnasoftDbConnection> {
    match open_database_connection(db_path) {
        Ok(mut c) => {
            trace!(
                "Checking if current existing database is compatible with Khulnasoft {}",
                khulnasoft_version
            );
            let db_version = c.query_row(
                "SELECT value FROM metadata WHERE key='KHULNASOFT_VERSION'",
                [],
                |row| {
                    let r: String = row.get(0)?;
                    Ok(r)
                },
            );
            let c = match db_version {
                Ok(Some(version)) if version == khulnasoft_version => {
                    trace!(
                        "Database is compatible with Khulnasoft {}",
                        khulnasoft_version
                    );
                    c
                }
                // If there is no metadata, it means that this database is new
                Err(s) if s.to_string().contains("metadata") => {
                    configure_database(&c)?;
                    create_metadata_table(&mut c, &khulnasoft_version)?;
                    c
                }
                reason => {
                    trace!("Incompatible database because: {:?}", reason);
                    trace!("Disconnecting from existing incompatible database");
                    c.close()?;
                    trace!("Removing existing incompatible database");
                    remove_file(db_path)?;

                    trace!("Initializing a new database");
                    initialize_db(khulnasoft_version, db_path)?
                }
            };

            Ok(c)
        }
        Err(reason) => {
            trace!(
                "Unable to connect to existing database because: {:?}",
                reason
            );
            trace!("Removing existing incompatible database");
            remove_file(db_path)?;

            trace!("Initializing a new database");
            initialize_db(khulnasoft_version, db_path)
        }
    }
}

fn create_metadata_table(
    c: &mut KhulnasoftDbConnection,
    khulnasoft_version: &str,
) -> anyhow::Result<()> {
    debug!("Creating table for metadata");
    c.transaction(|conn| {
        conn.execute(
            "CREATE TABLE metadata (
                key TEXT NOT NULL PRIMARY KEY,
                value TEXT NOT NULL
            )",
            [],
        )?;
        trace!("Recording Khulnasoft Version: {}", khulnasoft_version);
        conn.execute(
            "INSERT INTO metadata (key, value) VALUES ('KHULNASOFT_VERSION', ?)",
            [khulnasoft_version],
        )?;
        Ok(())
    })?;

    Ok(())
}

fn open_database_connection(db_path: &Path) -> anyhow::Result<KhulnasoftDbConnection> {
    let conn = Connection::open_with_flags(
        db_path,
        OpenFlags::SQLITE_OPEN_READ_WRITE
            | OpenFlags::SQLITE_OPEN_CREATE
            | OpenFlags::SQLITE_OPEN_URI
            | OpenFlags::SQLITE_OPEN_FULL_MUTEX,
    );

    conn.map_err(|e| anyhow::anyhow!("Error creating connection {:?}", e))
        .map(KhulnasoftDbConnection::new)
}

fn configure_database(connection: &KhulnasoftDbConnection) -> anyhow::Result<()> {
    connection
        .pragma_update(None, "journal_mode", "WAL")
        .map_err(|e| anyhow::anyhow!("Unable to set journal_mode: {:?}", e))?;
    connection
        .pragma_update(None, "synchronous", "NORMAL")
        .map_err(|e| anyhow::anyhow!("Unable to set synchronous: {:?}", e))?;
    connection
        .busy_handler(Some(|tries| tries <= 12))
        .map_err(|e| anyhow::anyhow!("Unable to set busy handler: {:?}", e))?;
    Ok(())
}

#[cfg(test)]
mod tests {
    use crate::native::logger::enable_logger;

    use super::*;

    #[test]
    fn initialize_db_creates_new_db() -> anyhow::Result<()> {
        let temp_dir = tempfile::tempdir()?;
        let db_path = temp_dir.path().join("test.db");

        let _ = initialize_db("1.0.0".to_string(), &db_path)?;

        let conn = Connection::open(&db_path)?;
        let version: String = conn.query_row(
            "SELECT value FROM metadata WHERE key='KHULNASOFT_VERSION'",
            [],
            |row| row.get(0),
        )?;

        assert_eq!(version, "1.0.0");
        Ok(())
    }

    #[test]
    fn initialize_db_reuses_compatible_db() -> anyhow::Result<()> {
        let temp_dir = tempfile::tempdir()?;
        let db_path = temp_dir.path().join("test.db");

        // Create initial db
        let _ = initialize_db("1.0.0".to_string(), &db_path)?;

        // Try to initialize again with same version
        let _ = initialize_db("1.0.0".to_string(), &db_path)?;

        let conn = Connection::open(&db_path)?;
        let version: String = conn.query_row(
            "SELECT value FROM metadata WHERE key='KHULNASOFT_VERSION'",
            [],
            |row| row.get(0),
        )?;

        assert_eq!(version, "1.0.0");
        Ok(())
    }

    #[test]
    fn initialize_db_recreates_incompatible_db() -> anyhow::Result<()> {
        enable_logger();
        let temp_dir = tempfile::tempdir()?;
        let db_path = temp_dir.path().join("test.db");
        //
        // Create initial db
        let _ = initialize_db("1.0.0".to_string(), &db_path)?;

        // Try to initialize with different version
        let conn = initialize_db("2.0.0".to_string(), &db_path)?;

        let version: Option<String> = conn.query_row(
            "SELECT value FROM metadata WHERE key='KHULNASOFT_VERSION'",
            [],
            |row| row.get(0),
        )?;

        assert_eq!(version.unwrap(), "2.0.0");
        Ok(())
    }
}
