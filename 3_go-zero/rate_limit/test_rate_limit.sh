
for i in {1..20};
do
  printf $i
  printf ' ->  '
  curl -i -X GET http://localhost:8889/greet?userId=1\&name=Smith
  printf '\n\n\n\n'
done