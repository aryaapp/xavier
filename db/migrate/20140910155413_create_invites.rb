class CreateInvites < ActiveRecord::Migration
  def change
    create_table :invites do |t|
      t.column :uuid, :uuid, default: "uuid_generate_v4()", :null => false

      t.string :message, :null => false
      t.string :status, :null => false

      t.integer :sender_id, :null => false
      t.string :sender_type, :null => false

      t.timestamps
    end
    
    change_column :invites, :uuid, :uuid, :null => false
    add_index :invites, :uuid
    add_reference :invites, :user, index: true    
  end
end
