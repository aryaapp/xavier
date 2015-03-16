class CreateNotes < ActiveRecord::Migration
  def change
    create_table :notes do |t|
      t.column :uuid, :uuid, default: "uuid_generate_v4()", :null => false
      t.string :content

      t.timestamps
    end
    
  	change_column :notes, :uuid, :uuid, :null => false
    add_index :notes, :uuid
    add_reference :notes, :app, index: true
    add_reference :notes, :user, index: true    
  end
end
