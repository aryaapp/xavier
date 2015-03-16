questionaire = Questionaire.find_by_id(1)
unless questionaire
  questionaire = Questionaire.create( 
  	uuid: "99dbdf90-327a-497b-8d1e-918d90a73642",
    title: "Default Questions",
    questions: [2,3,4,5,6]
  )
end

questionaire2 = Questionaire.find_by_id(2)
unless questionaire2
  questionaire2 = Questionaire.create( 
  	uuid: "c28978ad-61f4-4d11-aea3-7e9a563fe1ba", 
    title: "Yolo Questions",
    questions: [1,2,3,4,5,6,7]
  )
end