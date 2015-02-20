class CreateAnswers < ActiveRecord::Migration
  def change
    create_table :answers do |t|
      t.column :uuid, :uuid, null: false, default: "uuid_generate_v4()"

      t.jsonb :values
      t.boolean :answered, default: false, :null => false

      t.timestamps
    end
    change_column :answers, :uuid, :uuid, :null => false
    add_index :answers, :uuid

    add_reference :answers, :journal, index: true   
    add_reference :answers, :question, index: true   
  end
end
