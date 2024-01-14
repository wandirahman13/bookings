/var/www/mdc # cat tmp/test.sh
#!/bin/sh
GOOGLE_CHAT_WEBHOOK_URL="https://chat.googleapis.com/v1/spaces/AAAABI0BOFI/messages?key=AIzaSyDdI0hCZtE6vySjMm-WEfRq3CPzqKqqsHI&token=pk4hCTIqeLeZDbwXRs0eIkjv7HWDGx6EOBMu_ULOk10"

long_command="
a = 'console'
puts 'hello'
puts 'from'
puts a
"
OUTPUT=$(rails runner "$long_command")

curl --request POST "$GOOGLE_CHAT_WEBHOOK_URL" \
        --header 'Content-Type: application/json' \
        --data-raw "{'text':'$OUTPUT'}"
/var/www/mdc #