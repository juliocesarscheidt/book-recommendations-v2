# Evans GRPC Client

```bash

# install
curl \
  -L https://github.com/ktr0731/evans/releases/download/0.9.3/evans_linux_amd64.tar.gz \
  -o evans_linux_amd64.tar.gz
tar -xzvf evans_linux_amd64.tar.gz
rm -rf evans_linux_amd64.tar.gz
chmod +x evans && mv evans /usr/local/bin/



#################### User Service ####################
# CreateUser
echo '{"userRequest": {"name": "julio", "surname": "cesar", "email": "juliocesar@mail", "phone": ""}}' | evans -r cli call --host localhost github.com.juliocesarscheidt.usersmicroservice.UserService.CreateUser

# GetUser
echo '{"uuid": "17f7e9fcad70e012d21e0069"}' | evans -r cli call --host localhost github.com.juliocesarscheidt.usersmicroservice.UserService.GetUser

# UpdateUser
echo '{"uuid": "17f7e9fcad70e012d21e0069", "userRequest": {"name": "julio", "surname": "cesar", "email": "juliocesar@mail", "phone": "+554199999991"}}' | evans -r cli call --host localhost github.com.juliocesarscheidt.usersmicroservice.UserService.UpdateUser

# DeleteUser
echo '{"uuid": "17f7e9fcad70e012d21e0069"}' | evans -r cli call --host localhost github.com.juliocesarscheidt.usersmicroservice.UserService.DeleteUser

# ListUser
echo '{"page": 0, "size": 50}' | evans -r cli call --host localhost github.com.juliocesarscheidt.usersmicroservice.UserService.ListUser



#################### User Rate Service ####################
# UpsertUserRate
echo '{"userRateRequest": {"user_uuid": "17f7eae0a2005d96a50aa326", "book_uuid": "17f7eae0a2005d96a50aa328", "rate": 8.0}}' | evans -r cli call --host localhost github.com.juliocesarscheidt.usersmicroservice.UserRateService.UpsertUserRate

# GetUserRate
echo '{"user_uuid": "17f7eae0a2005d96a50aa326"}' | evans -r cli call --host localhost github.com.juliocesarscheidt.usersmicroservice.UserRateService.GetUserRate

# DeleteUserRate
echo '{"user_uuid": "17f7eae0a2005d96a50aa326"}' | evans -r cli call --host localhost github.com.juliocesarscheidt.usersmicroservice.UserRateService.DeleteUserRate

# ListUserRate
echo '{"page": 0, "size": 50}' | evans -r cli call --host localhost github.com.juliocesarscheidt.usersmicroservice.UserRateService.ListUserRate



# read eval print loop
evans --host localhost -r repl

> show package

> package [NAME]

> show service

> service [SERVICE]

> show message
> desc [MESSAGE]

> call [ENDPOINT]

```
