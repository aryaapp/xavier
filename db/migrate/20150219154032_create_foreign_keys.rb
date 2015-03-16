class CreateForeignKeys < ActiveRecord::Migration
  # Apps
  add_foreign_key :journals, :apps, on_update: :restrict, on_delete: :restrict
  add_foreign_key :notifications, :apps, on_update: :restrict, on_delete: :restrict
  add_foreign_key :users, :apps, on_update: :restrict, on_delete: :restrict


  # Journals
  add_foreign_key :answers, :journals, on_update: :cascade, on_delete: :cascade

  # Questionaires
  add_foreign_key :questionaires_users, :questionaires, on_update: :restrict, on_delete: :restrict

  # Questions
  add_foreign_key :answers, :questions, on_update: :restrict, on_delete: :restrict
  
  # Themes
  add_foreign_key :users, :themes, on_update: :restrict, on_delete: :restrict

  # Users
  add_foreign_key :devices, :users, on_update: :restrict, on_delete: :restrict
  add_foreign_key :invites, :users, on_update: :restrict, on_delete: :restrict
  add_foreign_key :journals, :users, on_update: :restrict, on_delete: :restrict
  add_foreign_key :keywords, :users, on_update: :restrict, on_delete: :restrict
  add_foreign_key :notes, :users, on_update: :restrict, on_delete: :restrict
  add_foreign_key :questionaires_users, :users, on_update: :restrict, on_delete: :restrict
end
