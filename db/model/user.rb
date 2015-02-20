class User < ActiveRecord::Base    
  has_secure_password

  has_one :access_token  
  belongs_to :theme

  has_many :devices
  has_many :invites
  has_many :journals
  has_many :keywords
  has_many :notes
  has_many :notifications

  has_many :relations,         foreign_key: "subject_id",  class_name: "Relation", dependent: :destroy
  has_many :reverse_relations, foreign_key: "observer_id", class_name: "Relation", dependent: :destroy

  has_many :subjects,  through: :relations, source: :subject
  has_many :observers, through: :reverse_relations, source: :observer

  has_and_belongs_to_many :questionaires
end