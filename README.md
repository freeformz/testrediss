```console
$ heroku create
$ heroku addons:create heroku-redis:premium-0
$ git push heroku master
$ heroku redis:wait
$ heroku run bash
$$ bin/redis.v5
$$ bin/redigo
```