class CreateQuestions < ActiveRecord::Migration
  def change
    create_table :questions do |t|
      t.column :uuid, :uuid, null: false, default: "uuid_generate_v4()"

      t.string :title, :null => false     
      t.string :description

      t.boolean :important, default: false, :null => false     
      t.boolean :autocompletes, default: false, :null => false     

      t.string :view, null: false
      t.string :processor, :null => false     
      t.column :user_data, :json

      t.timestamps
    end
     
    change_column :questions, :uuid, :uuid, :null => false
    add_index :questions, :uuid
  end
end
