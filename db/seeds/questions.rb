feelings = Question.find_by_id(1)
unless feelings
  feelings = Question.create( 
    uuid: "52b017a6-1b45-4544-b67f-0fdb0b4df4ae",
    title: "How are you feeling at the moment?", 
    description: "Move the sliders to indicate your feeling.",
    processor: "emotions",
    view: "sliders" 
  )
end
feelings.important = true
feelings.user_data = {
  "options" => [
    { "title" => "General feeling", "identifier" => "feeling", "min_value" => 0, "max_value" => 100, "default_value"=> 50, "steps" => 1 },
    { "title" => "Joy", "identifier" => "joy", "min_value" => 0, "max_value" => 100, "default_value"=> 50, "steps" => 1 },
    { "title" => "Trust", "identifier" => "trust", "min_value" => 0, "max_value" => 100, "default_value"=> 50, "steps" => 1 },
    { "title" => "Fear", "identifier" => "fear", "min_value" => 0, "max_value" => 100, "default_value"=> 50, "steps" => 1 },
    { "title" => "Surprise", "identifier" => "surprise", "min_value" => 0, "max_value" => 100, "default_value"=> 50, "steps" => 1 },
    { "title" => "Sadness", "identifier" => "sadness", "min_value" => 0, "max_value" => 100, "default_value"=> 50, "steps" => 1 },
    { "title" => "Lonely", "identifier" => "lonely", "min_value" => 0, "max_value" => 100, "default_value"=> 50, "steps" => 1 },
    { "title" => "Disgust", "identifier" => "disgust", "min_value" => 0, "max_value" => 100, "default_value"=> 50, "steps" => 1 },
    { "title" => "Anger", "identifier" => "anger", "min_value" => 0, "max_value" => 100, "default_value"=> 50, "steps" => 1 },
    { "title" => "Anticipation", "identifier" => "anticipation", "min_value" => 0, "max_value" => 100, "default_value"=> 50, "steps" => 1 }    
  ]
}
unless feelings.save 
  puts feelings.errors.full_messages
end
 

feeling = Question.find_by_id(2)
unless feeling
  feeling = Question.create(    
    uuid: "88396dd3-66da-4b61-a165-8bf415a1c5ea", 
    title: "How are you feeling at the moment?", 
    description: "Move the slider to indicate your feeling.",
    processor: "emotions",
    view: "sliders"  
  )
end
feeling.important = true
feeling.user_data = {
  "options" => [  
    { 
      "default_value"=> 50,
      "identifier" => "feeling", 
      "legend" => [ "VERY BAD", "AVERAGE", "VERY WELL" ],
      "max_value" => 100, 
      "min_value" => 0, 
      "steps" => 1,
      "title" => "Your feeling"
    }
  ]
}
feeling.save

circumstances = Question.find_by_id(3)
unless circumstances
  circumstances = Question.create( 
    uuid: "eeb7fb4a-fe2e-42df-b403-797bcdf6c7ae",
    title: "Describe the circumstances why you felt this way.",
    description: "Type in one or multiple items.",
    processor: "text-list", 
    view: "list" 
  )
end
circumstances.autocompletes = true
circumstances.save

body = Question.find_by_id(4)
unless body
  body = Question.create( 
    uuid: "2dfe8bc9-b8b8-4bea-97f9-a9f0e9eb52b8",
    title: "How does your body feel?",
    description: "Describe what you feel and where you feel it.",
    processor: "body", 
    view: "images" 
  )
end
body.user_data = {
  "image" => "body.svg",
  "type" => "svg",
  "images" => [ 
    { "frame" => { "x" => 69, "y"=> 0, "width" => 56, "height" => 49 }, "name" => "Head", "options" => [ "Headache", "Dizzy" ] }
  ]
}
body.save

thoughts = Question.find_by_id(5)
unless thoughts
  thoughts = Question.create(     
    uuid: "f87fea1b-32b1-4910-be3d-1db153713ec5",
    title: "What are your thoughts?",
    description: "You can type in what your thoughts are.",
    processor: "audio:audio, text:text" , 
    view: "text" 
  )
end
thoughts.user_data = { "placeholder" => "You can type in your thoughts here..." }
thoughts.save

reaction = Question.find_by_id(6)
unless reaction
  reaction = Question.create( 
    uuid: "81816d44-897c-4907-8311-95a1604fd459",
    title: "How did you react?",
    description: "Type in how you reacted to this situation.",
    processor: "text",
    view: "text" 
  )
end
reaction.user_data = { "placeholder" => "You can type in your reaction here..." }
reaction.save

explanation = Question.find_by_id(7)
unless explanation
  explanation = Question.create(
    uuid: "db271916-8581-40f8-a727-df8708348d1d",
    title: "What is it that you don't feel worse?",
    description: "Type in one or multiple items.",
    processor: "text-list",
    view: "list" 
  )
end
explanation.autocompletes = true
explanation.save
