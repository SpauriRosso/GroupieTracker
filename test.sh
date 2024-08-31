printf "\x1B[33mTesting base URL\x1B[0m\n"
wget -p --timestamping https://groupie.spyrisk.xyz
sleep 3
printf "\x1B[33mTesting base result URL, expected 405\x1B[0m\n"
wget -p --timestamping https://groupie.spyrisk.xyz/result
sleep 3
printf "\x1B[33mREMOVING TEST FILES\x1B[0m\n"
rm -r groupie.spyrisk.xyz
sleep 3
printf "\x1B[33mTEST COMPLETED\x1B[0m\n"

