class CreateUsers < ActiveRecord::Migration
  def change
    create_table :users do |t|
      t.column :uuid, :uuid, default: "uuid_generate_v4()", :null => false
      
      t.string :email, :null => false      
      t.string :password_digest, :null => false

      t.string  :gender, default: "male"      
      t.string  :fullname      
      t.string  :type_of_therapy
      
      t.boolean :public, default: false,  :null => false
      t.boolean :professional, default: false,  :null => false
      
      t.timestamps
    end
    change_column :users, :uuid, :uuid, :null => false

    add_index :users, :email, unique: true
    add_index :users, :uuid
  end
end
