class CreateQuestionairesUsers < ActiveRecord::Migration
  create_table :questionaires_users do |t|
    t.references :questionaire, :null => false
    t.references :user, :null => false
  end

  add_index(:questionaires_users, :questionaire_id)
  add_index(:questionaires_users, :user_id)    
end
