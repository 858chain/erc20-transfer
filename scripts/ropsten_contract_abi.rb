#!/usr/bin/env ruby

begin
  require 'httparty'
  require 'base64'
rescue LoadError  => e
  $stdout.puts e
end

$ua = 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36'


$contracts = {
  'tst' =>'0x722dd3F80BAC40c951b51BdD28Dd19d435762180',
}

def get_abi_and_convert_to_base64(name)
  path = "http://api-ropsten.etherscan.io/api?module=contract&action=getabi&address=#{$contracts[name]}&format=raw"
  ENV['http_proxy'] = ENV['https_proxy'] = nil
  resp = HTTParty.get(path, headers: { 'User-Agent' => $ua })
  Base64.encode64(resp.body)
end

$contracts.each_pair do |name, address|

  puts name
  hash = {"address": address, "abi": get_abi_and_convert_to_base64(name)}
  puts hash.to_json
  puts "-" * 100
end
