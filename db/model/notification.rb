class Notification < ActiveRecord::Base
  belongs_to :app
  belongs_to :user
  
end
