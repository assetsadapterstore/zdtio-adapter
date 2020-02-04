module github.com/assetsadapterstore/zdtio-adapter

go 1.12

require (
	github.com/astaxie/beego v1.12.0
	github.com/blocktree/eosio-adapter v1.4.0
	github.com/blocktree/go-owcdrivers v1.2.0
	github.com/blocktree/openwallet v1.7.0
	github.com/eoscanada/eos-go v0.8.16
)

replace github.com/eoscanada/eos-go => github.com/blocktree/eos-go v0.8.13-blocktree
