hello there
this an api that can give you a unreal environment for setting and getting movie and
showing all movies you set.
you can set option for the movies as:
  fullname
  country
  year
  imdb

for starting you do in two ways:  1-docker  2-git
1- you can enter following command into the terminal or cmd for getting the image:
  # docker pull arash2007mahdavi/digimovie:1.0
  # docker run -d --name digimovie_container -p 8080:1386 arash2007mahdavi/digimovie:1.0
  now its run on your "http://localhost:8080"
2- and you can get the full projects files from github:
  # git clone https://github.com/arash2007mahdavi/digimovie.git
  now you have full project. you should just run the "main.go" file from src folder
  now its run on your "http://localhost:8080" (you must have golang on your system)

command: (api-key: 09945092968)
1- digimovie
2- digimovie/version
3- digimovie/set/movie: need: (fullname, country, year, ibdm, METHOD:post)
4- digimovie/get/movie: need: (fullname, METHOD:get)
5- digimovie/movies/list: need: (METHOD:get)
