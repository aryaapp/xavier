class CreateThemes < ActiveRecord::Migration
  def change
    create_table :themes do |t|
      t.column :uuid, :uuid, default: "uuid_generate_v4()", :null => false
      t.string :color
      t.string :wallpaper
      t.string :type

      t.timestamps
    end
    change_column :themes, :uuid, :uuid, :null => false
    add_index :themes, :uuid
    add_reference :users, :theme, index: true
  end
end
