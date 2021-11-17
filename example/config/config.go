//nolint
package config

const Baseurl = "baseurl"

type build struct {
	Nojsconfiginassets string

	Useresourcecachewhen string

	Writestats string
}

var Build = build{

	Nojsconfiginassets: "build.nojsconfiginassets",

	Useresourcecachewhen: "build.useresourcecachewhen",

	Writestats: "build.writestats",
}

type caches struct {
	Assets cachesAssets

	Getcsv cachesGetcsv

	Getjson cachesGetjson

	Images cachesImages

	Modules cachesModules
}

var Caches = caches{

	Assets: CachesAssets,

	Getcsv: CachesGetcsv,

	Getjson: CachesGetjson,

	Images: CachesImages,

	Modules: CachesModules,
}

type cachesAssets struct {
	Dir string

	Maxage string
}

var CachesAssets = cachesAssets{

	Dir: "caches.assets.dir",

	Maxage: "caches.assets.maxage",
}

type cachesGetcsv struct {
	Dir string

	Maxage string
}

var CachesGetcsv = cachesGetcsv{

	Dir: "caches.getcsv.dir",

	Maxage: "caches.getcsv.maxage",
}

type cachesGetjson struct {
	Dir string

	Maxage string
}

var CachesGetjson = cachesGetjson{

	Dir: "caches.getjson.dir",

	Maxage: "caches.getjson.maxage",
}

type cachesImages struct {
	Dir string

	Maxage string
}

var CachesImages = cachesImages{

	Dir: "caches.images.dir",

	Maxage: "caches.images.maxage",
}

type cachesModules struct {
	Dir string

	Maxage string
}

var CachesModules = cachesModules{

	Dir: "caches.modules.dir",

	Maxage: "caches.modules.maxage",
}

const Footnotereturnlinkcontents = "footnotereturnlinkcontents"

type params struct {
	Authorname string

	Debug string

	Githubuser string

	Listoffoo string

	Sidebarrecentlimit string

	Subtitle string
}

var Params = params{

	Authorname: "params.authorname",

	Debug: "params.debug",

	Githubuser: "params.githubuser",

	Listoffoo: "params.listoffoo",

	Sidebarrecentlimit: "params.sidebarrecentlimit",

	Subtitle: "params.subtitle",
}

type permalinks struct {
	Posts string
}

var Permalinks = permalinks{

	Posts: "permalinks.posts",
}

const Title = "title"
