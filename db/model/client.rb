class Client < ActiveRecord::Base  
  has_many :access_tokens
  has_and_belongs_to_many :questions
end

scopes = [ 
  "connections",
  "devices", 
  "invites", 
  "journals",
  "keywords",
  "notes", 
  "notifications", 
  "profile", 
  "users#search",
  "themes#update"
]

grant_types =  [ "authorization_code", "password", "refresh_token" ]

ios = Client.find_or_create_by( 
  uuid: "668ac08d-82b4-42a6-943b-2f6ca2c38258",
  name: "Arya iOS App", 
  url: "http://aryaapp.co",
  secret: "757d831680161c4b351c444b0719d874445115a0e7dc2673e53817879c44607f"
)
ios.grant_types = grant_types
ios.permitted_scopes = scopes
ios.save

android = Client.find_or_create_by( 
  uuid: "e782c0c0-cf2c-4720-929f-bdcd314028f7",
  name: "Arya Android App", 
  url: "http://aryaapp.co",
  secret: "b4c081b1561238cf9b5e22e9c4deb67b10d39e8ebe748e2f334c7afeba6875bb"
)
android.grant_types = grant_types
android.permitted_scopes = scopes
android.save