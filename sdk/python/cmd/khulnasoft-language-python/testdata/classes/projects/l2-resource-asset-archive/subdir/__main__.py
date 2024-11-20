import khulnasoft
import khulnasoft_asset_archive as asset_archive

ass = asset_archive.AssetResource("ass", value=khulnasoft.FileAsset("../test.txt"))
arc = asset_archive.ArchiveResource("arc", value=khulnasoft.FileArchive("../archive.tar"))
dir = asset_archive.ArchiveResource("dir", value=khulnasoft.FileArchive("../folder"))
assarc = asset_archive.ArchiveResource("assarc", value=khulnasoft.AssetArchive({
    "string": khulnasoft.StringAsset("file contents"),
    "file": khulnasoft.FileAsset("../test.txt"),
    "folder": khulnasoft.FileArchive("../folder"),
    "archive": khulnasoft.FileArchive("../archive.tar"),
}))
remoteass = asset_archive.AssetResource("remoteass", value=khulnasoft.RemoteAsset("https://raw.githubusercontent.com/khulnasoft/khulnasoft/master/cmd/khulnasoft-test-language/testdata/l2-resource-asset-archive/test.txt"))
