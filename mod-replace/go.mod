module example.com/mod-replace

go 1.15

require example.com/test v0.0.0
replace example.com/test v0.0.0 => ../mod-test
