use parking_lot::{Condvar, Mutex, MutexGuard};

pub struct KhulnasoftMutex<T>(Mutex<T>);

impl<T> KhulnasoftMutex<T> {
    pub fn new(value: T) -> Self {
        Self(Mutex::new(value))
    }
    pub fn lock(&self) -> anyhow::Result<MutexGuard<'_, T>> {
        Ok(self.0.lock())
    }
}

pub struct KhulnasoftCondvar(Condvar);

impl KhulnasoftCondvar {
    pub fn new() -> Self {
        Self(Condvar::new())
    }

    pub fn wait<'a, T, F>(
        &'a self,
        mut guard: MutexGuard<'a, T>,
        condition: F,
    ) -> anyhow::Result<MutexGuard<'a, T>>
    where
        F: Fn(&MutexGuard<'a, T>) -> bool,
    {
        if condition(&guard) {
            self.0.wait(&mut guard);
        }
        Ok(guard)
    }

    pub fn notify_all(&self) {
        self.0.notify_all();
    }
}
