module demo-user

go 1.14

replace demo-transaction/proto v0.0.0 => ./proto

require (
	demo-transaction/proto v0.0.0
	github.com/asaskevich/govalidator v0.0.0-20200819183940-29e1ff8eb0bb
	github.com/go-ozzo/ozzo-validation/v4 v4.2.2
	github.com/labstack/echo/v4 v4.1.17
	github.com/lqhoang99/cashbag-proto v0.0.0-20200903070255-4f620ce0406e // indirect
	github.com/samuel/go-zookeeper v0.0.0-20200724154423-2164a8ac840e
	github.com/stretchr/testify v1.4.0
	go.mongodb.org/mongo-driver v1.4.1
	google.golang.org/grpc v1.32.0
)
