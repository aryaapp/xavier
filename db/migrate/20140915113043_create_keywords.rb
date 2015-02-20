class CreateKeywords < ActiveRecord::Migration
  def change
    create_table :keywords do |t|
      t.column :uuid, :uuid, default: "uuid_generate_v4()", :null => false
      
      t.string  :name, :null => false
      t.integer :count, default: 1, :null => false
      t.integer :relevance, default: 0, :null => false

      t.timestamps
    end

    change_column :keywords, :uuid, :uuid, :null => false

    add_reference :keywords, :question, index: true   
    add_reference :keywords, :user, index: true   
    
    add_index :keywords, :uuid
    add_index :keywords, :name
    add_index(:keywords, [:name, :question_id, :user_id], :unique => true)
  end
end
