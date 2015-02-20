menno = User.find_by(email: "ome.menno@gmail.com")
menno = User.create( email: "ome.menno@gmail.com", password: "helloworld", password_confirmation: "helloworld" ) unless menno

purcy = User.find_by(email: "purcymarte@gmail.com")
purcy = User.create( email: "purcymarte@gmail.com", password: "helloworld", password_confirmation: "helloworld" ) unless purcy

mark = User.find_by(email: "mark.eibes@googlemail.com")
mark = User.create( email: "mark.eibes@googlemail.com", password: "helloworld", password_confirmation: "helloworld" ) unless mark

yoloboy = User.find_by(email: "yoloboy@yolomail.com")
yoloboy = User.create( email: "yoloboy@yolomail.com", password: "helloworld", password_confirmation: "helloworld" ) unless yoloboy