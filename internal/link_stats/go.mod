module shortlink/internal/link_stats

go 1.23.0

replace (
	shortlink/internal/base => ../base
	shortlink/internal/link => ../link // 只应该依赖 link 模块的 domain 层
)

require (
	github.com/gofiber/fiber/v2 v2.52.5
	github.com/jinzhu/copier v0.4.0
	github.com/redis/go-redis/v9 v9.6.1
	gorm.io/gorm v1.25.11
	shortlink/internal/base v0.0.0-00010101000000-000000000000
	shortlink/internal/link v0.0.0-00010101000000-000000000000
)

require (
	github.com/PuerkitoBio/goquery v1.9.2 // indirect
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/andybalholm/cascadia v1.3.2 // indirect
	github.com/bsm/redislock v0.9.4 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/klauspost/compress v1.17.2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/mssola/user_agent v0.6.0 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.51.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)
