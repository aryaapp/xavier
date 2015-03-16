Dir[File.expand_path("model/**/*.rb", File.dirname(__FILE__))].each { |rb| require rb }

%w{ apps questionaires questions themes users }.each do |part|
  require File.expand_path(File.dirname(__FILE__))+"/seeds/#{part}.rb"
end

questionaire = Questionaire.find_by_id(1)
yolo_questionaire = Questionaire.find_by_id(2)

menno = User.find_by(email: "ome.menno@gmail.com")
menno.theme = Theme.find_by(color: "ff0066", wallpaper: "TestBackground.jpg")
menno.questionaires = [ questionaire ]
menno.save

purcy = User.find_by(email: "purcymarte@gmail.com")
purcy.theme = Theme.find_by(color: "379cd3", wallpaper: "TestBackground2.jpg")
purcy.questionaires =  [ questionaire ]
purcy.save

mark = User.find_by(email: "mark.eibes@googlemail.com")
mark.theme = Theme.find_by(color: "038c0d", wallpaper: "TestBackground3.jpg")
mark.questionaires =  [ questionaire ]
mark.save

yoloboy = User.find_by(email: "yoloboy@yolomail.com")
yoloboy.theme = Theme.find_by(color: "ff0066", wallpaper: "TestBackground.jpg")
yoloboy.questionaires =  [ yolo_questionaire ]
yoloboy.public = true
yoloboy.professional = true
yoloboy.save