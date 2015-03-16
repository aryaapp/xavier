menno = User.find_by( email: "ome.menno@gmail.com")
menno = User.create( 
	uuid: "cffb820d-e8a9-48bb-9e9b-1ebd65b82d2", 
	email: "ome.menno@gmail.com", 
	password: "helloworld", 
	password_confirmation: "helloworld",
	app_id: 1
) unless menno

purcy = User.find_by(email: "purcymarte@gmail.com")
purcy = User.create( 
	uuid: "13cf5211-05db-4d21-a328-767237309e35",
	email: "purcymarte@gmail.com", 
	password: "helloworld", 
	password_confirmation: "helloworld",
	app_id: 2
) unless purcy

mark = User.find_by(email: "mark.eibes@googlemail.com")
mark = User.create( 
	uuid: "53d8539e-e913-4467-a875-af5a7e7cf84c", 
	email: "mark.eibes@googlemail.com", 
	password: "helloworld", 
	password_confirmation: "helloworld",
	app_id: 2
) unless mark

yoloboy = User.find_by(email: "yoloboy@yolomail.com")
yoloboy = User.create( 
	uuid: "0195817f-93b4-4124-a800-c6d8badc4330", 
	email: "yoloboy@yolomail.com", 
	password: "helloworld", 
	password_confirmation: "helloworld",
	app_id: 1
) unless yoloboy