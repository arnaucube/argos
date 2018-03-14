http://arnaucode.com/argos/

# argos [![Go Report Card](https://goreportcard.com/badge/github.com/arnaucode/argos)](https://goreportcard.com/report/github.com/arnaucode/argos)
Open source twitter entropic toolkit. The main idea is to analyze, clean and generate public data, to add more entropy to the network.

Be half bot and half human, a new generation of cyborgs

The user analyzer features are insipired on the Python tweets_analyzer by x0rz https://github.com/x0rz/tweets_analyzer

#### [Disclamer]
This software was made to learn the Go syntax, it works, but the implementation is not in the optimal way. I don't have time to maintain it.

#### Argos Panoptes
https://en.wikipedia.org/wiki/Argus_Panoptes

https://en.wikipedia.org/wiki/Panopticon


![Argos](https://raw.githubusercontent.com/arnaucode/argos/master/argos.jpg "argos")


#### Current features
1. User analyzer
    - Word count
    - Weekly activity distribution
    - Daily activity distribution
    - Devices used
    - Hashtags most used count
    - Interactions with other users count
2. Unfollow all
3. Random follow
    - selects n number of accounts to follow, and follows n random accounts
4. Delete Tweets
5. Delete Favs (Likes)
6. Random tweet
    - post a tweet with content from a selected account
7. Tweet analyzer
    - analyze the retweets and detect Bots

![screen](https://raw.githubusercontent.com/arnaucode/argos/master/screen3.png "screen")

![screen](https://raw.githubusercontent.com/arnaucode/argos/master/screen2.png "screen")

needs a twitterConfig.json file on the /build folder with the content:
```
{
    "consumer_key": "xxxxxxxxxxxxxxxx",
    "consumer_secret": "xxxxxxxxxxxxxxxx",
    "access_token_key": "xxxxxxxxxxxxxxxx",
    "access_token_secret": "xxxxxxxxxxxxxxxx"
}

```

to run it:
- go to the /build folder
- open terminal
- execute the script with:
```
./argos
```
- follow the instructions that appears on the terminal
