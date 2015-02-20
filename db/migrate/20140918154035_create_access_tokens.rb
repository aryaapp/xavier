class CreateAccessTokens < ActiveRecord::Migration
  def change
    create_table :access_tokens do |t|
      t.column :uuid, :uuid, null: false, default: "uuid_generate_v4()"

      t.string :token, null: false
      t.string :scopes, array: true, default: []

      t.timestamps
    end
        
    change_column :access_tokens, :uuid, :uuid, :null => false

    add_reference :access_tokens, :client, index: true 
    add_reference :access_tokens, :user, index: true  
    
    add_index :access_tokens, :uuid
    add_index(:access_tokens, [:token, :client_id], :unique => true)
  end
end
