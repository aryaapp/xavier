questionaire = Questionaire.find_by_id(1)
unless questionaire
  questionaire = Questionaire.create( 
    title: "Default Questions",
    questions: [2,3,4,5,6]
  )
end

questionaire2 = Questionaire.find_by_id(2)
unless questionaire2
  questionaire2 = Questionaire.create( 
    title: "Yolo Questions",
    questions: [1,2,3,4,5,6,7]
  )
end