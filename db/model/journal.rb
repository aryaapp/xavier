class Journal < ActiveRecord::Base
  belongs_to :app
  belongs_to :user
  
  has_many :answers
end
