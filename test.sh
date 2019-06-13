#!/bin/bash
#
# simple tests for the endpoints using CURL
# asuming there is user Bob having some relevant data store in his account
RED="\033[0;31m'"
GREEN="\033[0;32m"
WHITE="\033[0m"

function cecho
{
  if [[ $1 == "r" ]]; then
    echo -e "${RED}"
    echo -n "$2"
  elif [[ $1 == "g" ]]; then
    echo -e "${GREEN}"
    echo -n "$2"
  else
    echo -e "${WHITE}"
    echo -n "$2"
  fi
 echo -e "${WHITE}"
 echo -e "___________________________________________________________"
}

printf "\033c"
#
## Test successful 401 responses from calling endpoints without an API key
#
cecho w "TEST 401 response"
echo "Check 401 on /social-networks"
response=$(curl --write-out %{http_code} --silent --output /dev/null  "http://127.0.0.1:2222/v1/social-networks")
if [[ $response == 401 ]];then cecho g "test pass"; else cecho r "test failed expected response 401 actual response: $response"; fi


echo "Check 401 on /resolve/social/{network}/{username}"
response=$(curl --write-out %{http_code} --silent --output /dev/null  "http://127.0.0.1:2222/v1/resolve/social/email/bob@example.com")
if [[ $response == 401 ]];then cecho g "test pass"; else cecho r "test failed expected response 401 actual response: $response"; fi

echo "Check 401 on /resolve/qrcode/{network}/{username}"
response=$(curl --write-out %{http_code} --silent --output /dev/null  "http://127.0.0.1:2222/v1/resolve/qrcode/email/bob@example.com")
if [[ $response == 401 ]];then cecho g "test pass"; else cecho r "test failed expected response 401 actual response: $response"; fi

echo "Check 401 on /resolve/qrcode/html/{network}/{username}"
response=$(curl --write-out %{http_code} --silent --output /dev/null  "http://127.0.0.1:2222/v1/resolve/qrcode/html/email/bob@example.com")
if [[ $response == 401 ]];then cecho g "test pass"; else cecho r "test failed expected response 401 actual response: $response"; fi

echo "Check 401 on /resolve/xrplaccount/{xrplaccount}/{tag}"
response=$(curl --write-out %{http_code} --silent --output /dev/null  "http://127.0.0.1:2222/v1/resolve/xrplaccount/rAsDfasdfADFssadfADFasdfDFsdv/55")
if [[ $response == 401 ]];then cecho g "test pass"; else cecho r "test failed expected response 401 actual response: $response"; fi

echo "Check 401 on /resolve/user/{slug}"
response=$(curl --write-out %{http_code} --silent --output /dev/null  "http://127.0.0.1:2222/v1/resolve/user/bob")
if [[ $response == 401 ]];then cecho g "test pass"; else cecho r "test failed expected response 401 actual response: $response"; fi
#
## Test successful responses from calling endpoints without an API key
#
cecho w  "TEST 200 response"

echo "Check 200 on /social-networks"
response=$(curl --write-out %{http_code} --silent --output /dev/null -H "XRPLNS-KEY: 8715785b8139ba56eae548e98d24ef65214b09c4" "http://127.0.0.1:2222/v1/social-networks")
if [[ $response == 200 ]];then cecho g "test pass"; else cecho r "test failed expected response 200 actual response: $response"; fi

echo "Check 200 on /resolve/social/{network}/{username}"
response=$(curl --write-out %{http_code} --silent --output /dev/null -H "XRPLNS-KEY: 8715785b8139ba56eae548e98d24ef65214b09c4" "http://127.0.0.1:2222/v1/resolve/social/email/bob@example.com")
if [[ $response == 200 ]];then cecho g "test pass"; else cecho r "test failed expected response 200 actual response: $response"; fi

echo "Check 200 on /resolve/qrcode/{network}/{username}"
response=$(curl --write-out %{http_code} --silent --output /dev/null -H "XRPLNS-KEY: 8715785b8139ba56eae548e98d24ef65214b09c4" "http://127.0.0.1:2222/v1/resolve/qrcode/email/bob@example.com")
if [[ $response == 200 ]];then cecho g "test pass"; else cecho r "test failed expected response 200 actual response: $response"; fi

echo "Check 200 on /resolve/qrcode/html/{network}/{username}"
response=$(curl --write-out %{http_code} --silent --output /dev/null -H "XRPLNS-KEY: 8715785b8139ba56eae548e98d24ef65214b09c4" "http://127.0.0.1:2222/v1/resolve/qrcode/html/email/bob@example.com")
if [[ $response == 200 ]];then cecho g "test pass"; else echo "test failed expected response 200 actual response: $response"; fi

echo "Check 200 on /resolve/xrplaccount/{xrplaccount}/{tag}"
response=$(curl --write-out %{http_code} --silent --output /dev/null -H "XRPLNS-KEY: 8715785b8139ba56eae548e98d24ef65214b09c4" "http://127.0.0.1:2222/v1/resolve/xrplaccount/rAsDfasdfADFssadfADFasdfDFsdv/55")
if [[ $response == 200 ]];then cecho g "test pass"; else cecho r "test failed expected response 200 actual response: $response"; fi

echo "Check 200 on /resolve/user/{slug}"
response=$(curl --write-out %{http_code} --silent --output /dev/null -H "XRPLNS-KEY: 8715785b8139ba56eae548e98d24ef65214b09c4" "http://127.0.0.1:2222/v1/resolve/user/bob")
if [[ $response == 200 ]];then cecho g "test pass"; else cecho r "test failed expected response 200 actual response: $response"; fi
#
## Test successful responses from calling endpoints without an API key
#
cecho w  "TEST actual response data"
echo "Check 200 on /social-networks"
response=$(curl --silent -H "XRPLNS-KEY: 8715785b8139ba56eae548e98d24ef65214b09c4" "http://127.0.0.1:2222/v1/social-networks")
if [[ "${response[@]}" =~ "email" ]];then cecho g "test pass"; else cecho r "test failed expected response 'email website' actual response: $response"; fi


echo "Check 200 on /resolve/social/{network}/{username}"
response=$(curl --silent -H "XRPLNS-KEY: 8715785b8139ba56eae548e98d24ef65214b09c4" "http://127.0.0.1:2222/v1/resolve/social/email/bob@example.com")
if [[ "${response[@]}" =~ "bob" ]];then cecho g "test pass"; else cecho r "test failed expected response 200 actual response: $response"; fi

echo "Check 200 on /resolve/qrcode/{network}/{username}"
response=$(curl --silent -H "XRPLNS-KEY: 8715785b8139ba56eae548e98d24ef65214b09c4" "http://127.0.0.1:2222/v1/resolve/qrcode/email/bob@example.com")
if [[ "${response[@]}" =~ "qrcode" ]];then cecho g "test pass"; else cecho r "test failed expected response 200 actual response: $response"; fi

echo "Check 200 on /resolve/qrcode/html/{network}/{username}"
response=$(curl --silent -H "XRPLNS-KEY: 8715785b8139ba56eae548e98d24ef65214b09c4" "http://127.0.0.1:2222/v1/resolve/qrcode/html/email/bob@example.com")
if [[ "${response[@]}" =~ "qrcode" ]];then cecho g "test pass"; else cecho r "test failed expected response 200 actual response: $response"; fi

echo "Check 200 on /resolve/xrplaccount/{xrplaccount}/{tag}"
response=$(curl --silent -H "XRPLNS-KEY: 8715785b8139ba56eae548e98d24ef65214b09c4" "http://127.0.0.1:2222/v1/resolve/xrplaccount/rAsDfasdfADFssadfADFasdfDFsdv/55")
if [[ "${response[@]}" =~ "bob" ]];then cecho g "test pass"; else cecho r "test failed expected response 200 actual response: $response"; fi

echo "Check 200 on /resolve/user/{slug}"
response=$(curl --silent -H "XRPLNS-KEY: 8715785b8139ba56eae548e98d24ef65214b09c4" "http://127.0.0.1:2222/v1/resolve/user/bob")
if [[ "${response[@]}" =~ "bob" ]];then cecho g "test pass"; else cecho r "test failed expected response 200 actual response: $response"; fi
