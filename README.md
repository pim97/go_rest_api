**VueJS & Go REST Api combined**
1. The VueJS will send a HTTP GET request to the GO REST Api with Axios
2. The Go API will fetch the details for the specific IMDB ID and put it in the SQLLite database
3. The VueJS application will send a HTTP GET request to fetch all the current movies added to the database
4. I demonstrated amazing reactive programming :)

**How to run**

1. unzip the results.zip (most likely already done)

2. get the following repo's
go get github.com/gin-gonic/gin
go get github.com/jinzhu/gorm
go get github.com/jinzhu/gorm/dialects/sqlite

3. go run movie.go

4. server setup at port 1991

5. database sqlite example.db is already included

7. npm run serve to run the VueJS application

8. Add an IMDB ID to the input field and click the button

![Test Image 3](https://i.gyazo.com/ea66cdb622508860dedcc13cedad5762.gif)
