- error when getting user from user following list, and this user don't have users following

```
optionFollowRandom.go
line 33
```
just need to check if the user have following users, and if not, get another user to follow
