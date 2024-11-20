import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as asset_archive from "@khulnasoft/asset-archive";

const ass = new asset_archive.AssetResource("ass", {value: new khulnasoft.asset.FileAsset("../test.txt")});
const arc = new asset_archive.ArchiveResource("arc", {value: new khulnasoft.asset.FileArchive("../archive.tar")});
const dir = new asset_archive.ArchiveResource("dir", {value: new khulnasoft.asset.FileArchive("../folder")});
const assarc = new asset_archive.ArchiveResource("assarc", {value: new khulnasoft.asset.AssetArchive({
    string: new khulnasoft.asset.StringAsset("file contents"),
    file: new khulnasoft.asset.FileAsset("../test.txt"),
    folder: new khulnasoft.asset.FileArchive("../folder"),
    archive: new khulnasoft.asset.FileArchive("../archive.tar"),
})});
const remoteass = new asset_archive.AssetResource("remoteass", {value: new khulnasoft.asset.RemoteAsset("https://raw.githubusercontent.com/khulnasoft/khulnasoft/master/cmd/khulnasoft-test-language/testdata/l2-resource-asset-archive/test.txt")});
