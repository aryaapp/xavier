# encoding: UTF-8
# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# Note that this schema.rb definition is the authoritative source for your
# database schema. If you need to create the application database on another
# system, you should be using db:schema:load, not running all the migrations
# from scratch. The latter is a flawed and unsustainable approach (the more migrations
# you'll amass, the slower it'll run and the greater likelihood for issues).
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema.define(version: 20150219154032) do

  # These are extensions that must be enabled in order to support this database
  enable_extension "plpgsql"
  enable_extension "uuid-ossp"
  enable_extension "hstore"

  create_table "access_tokens", force: :cascade do |t|
    t.uuid     "uuid",       default: "uuid_generate_v4()", null: false
    t.string   "token",                                     null: false
    t.string   "scopes",     default: [],                                array: true
    t.datetime "created_at"
    t.datetime "updated_at"
    t.integer  "client_id"
    t.integer  "user_id"
  end

  add_index "access_tokens", ["client_id"], name: "index_access_tokens_on_client_id", using: :btree
  add_index "access_tokens", ["token", "client_id"], name: "index_access_tokens_on_token_and_client_id", unique: true, using: :btree
  add_index "access_tokens", ["user_id"], name: "index_access_tokens_on_user_id", using: :btree
  add_index "access_tokens", ["uuid"], name: "index_access_tokens_on_uuid", using: :btree

  create_table "answers", force: :cascade do |t|
    t.uuid     "uuid",        default: "uuid_generate_v4()", null: false
    t.jsonb    "values"
    t.boolean  "answered",    default: false,                null: false
    t.datetime "created_at"
    t.datetime "updated_at"
    t.integer  "journal_id"
    t.integer  "question_id"
  end

  add_index "answers", ["journal_id"], name: "index_answers_on_journal_id", using: :btree
  add_index "answers", ["question_id"], name: "index_answers_on_question_id", using: :btree
  add_index "answers", ["uuid"], name: "index_answers_on_uuid", using: :btree

  create_table "clients", force: :cascade do |t|
    t.uuid     "uuid",             default: "uuid_generate_v4()", null: false
    t.string   "name",                                            null: false
    t.string   "url",                                             null: false
    t.string   "secret",                                          null: false
    t.string   "grant_types",      default: [],                   null: false, array: true
    t.string   "permitted_scopes", default: [],                   null: false, array: true
    t.datetime "created_at"
    t.datetime "updated_at"
  end

  add_index "clients", ["uuid"], name: "index_clients_on_uuid", using: :btree

  create_table "devices", force: :cascade do |t|
    t.string   "token",       null: false
    t.string   "environment", null: false
    t.string   "name",        null: false
    t.string   "model",       null: false
    t.string   "os",          null: false
    t.string   "os_version",  null: false
    t.string   "app_version", null: false
    t.datetime "created_at"
    t.datetime "updated_at"
    t.integer  "user_id"
  end

  add_index "devices", ["token"], name: "index_devices_on_token", unique: true, using: :btree
  add_index "devices", ["user_id"], name: "index_devices_on_user_id", using: :btree

  create_table "invites", force: :cascade do |t|
    t.uuid     "uuid",        default: "uuid_generate_v4()", null: false
    t.string   "message",                                    null: false
    t.string   "status",                                     null: false
    t.integer  "sender_id",                                  null: false
    t.string   "sender_type",                                null: false
    t.datetime "created_at"
    t.datetime "updated_at"
    t.integer  "user_id"
  end

  add_index "invites", ["user_id"], name: "index_invites_on_user_id", using: :btree
  add_index "invites", ["uuid"], name: "index_invites_on_uuid", using: :btree

  create_table "journals", force: :cascade do |t|
    t.uuid     "uuid",       default: "uuid_generate_v4()", null: false
    t.float    "feeling",    default: -1.0,                 null: false
    t.jsonb    "questions",                                 null: false
    t.datetime "created_at"
    t.datetime "updated_at"
    t.integer  "client_id"
    t.integer  "user_id"
  end

  add_index "journals", ["client_id"], name: "index_journals_on_client_id", using: :btree
  add_index "journals", ["user_id"], name: "index_journals_on_user_id", using: :btree
  add_index "journals", ["uuid"], name: "index_journals_on_uuid", using: :btree

  create_table "keywords", force: :cascade do |t|
    t.uuid     "uuid",        default: "uuid_generate_v4()", null: false
    t.string   "name",                                       null: false
    t.integer  "count",       default: 1,                    null: false
    t.integer  "relevance",   default: 0,                    null: false
    t.datetime "created_at"
    t.datetime "updated_at"
    t.integer  "question_id"
    t.integer  "user_id"
  end

  add_index "keywords", ["name", "question_id", "user_id"], name: "index_keywords_on_name_and_question_id_and_user_id", unique: true, using: :btree
  add_index "keywords", ["name"], name: "index_keywords_on_name", using: :btree
  add_index "keywords", ["question_id"], name: "index_keywords_on_question_id", using: :btree
  add_index "keywords", ["user_id"], name: "index_keywords_on_user_id", using: :btree
  add_index "keywords", ["uuid"], name: "index_keywords_on_uuid", using: :btree

  create_table "notes", force: :cascade do |t|
    t.uuid     "uuid",       default: "uuid_generate_v4()", null: false
    t.string   "content"
    t.datetime "created_at"
    t.datetime "updated_at"
    t.integer  "user_id"
  end

  add_index "notes", ["user_id"], name: "index_notes_on_user_id", using: :btree
  add_index "notes", ["uuid"], name: "index_notes_on_uuid", using: :btree

  create_table "notifications", force: :cascade do |t|
    t.uuid     "uuid",        default: "uuid_generate_v4()", null: false
    t.string   "message",                                    null: false
    t.boolean  "read",        default: false
    t.integer  "object_id"
    t.string   "object_type"
    t.string   "object_uri"
    t.datetime "created_at"
    t.datetime "updated_at"
    t.integer  "user_id"
  end

  add_index "notifications", ["user_id"], name: "index_notifications_on_user_id", using: :btree
  add_index "notifications", ["uuid"], name: "index_notifications_on_uuid", using: :btree

  create_table "questionaires", force: :cascade do |t|
    t.uuid     "uuid",       default: "uuid_generate_v4()", null: false
    t.string   "title",                                     null: false
    t.integer  "questions",  default: [],                   null: false, array: true
    t.datetime "created_at"
    t.datetime "updated_at"
  end

  add_index "questionaires", ["uuid"], name: "index_questionaires_on_uuid", using: :btree

  create_table "questionaires_users", force: :cascade do |t|
    t.integer "questionaire_id", null: false
    t.integer "user_id",         null: false
  end

  add_index "questionaires_users", ["questionaire_id"], name: "index_questionaires_users_on_questionaire_id", using: :btree
  add_index "questionaires_users", ["user_id"], name: "index_questionaires_users_on_user_id", using: :btree

  create_table "questions", force: :cascade do |t|
    t.uuid     "uuid",          default: "uuid_generate_v4()", null: false
    t.string   "title",                                        null: false
    t.string   "description"
    t.boolean  "important",     default: false,                null: false
    t.boolean  "autocompletes", default: false,                null: false
    t.string   "view",                                         null: false
    t.string   "processor",                                    null: false
    t.json     "user_data"
    t.datetime "created_at"
    t.datetime "updated_at"
  end

  add_index "questions", ["uuid"], name: "index_questions_on_uuid", using: :btree

  create_table "relations", force: :cascade do |t|
    t.integer  "observer_id", null: false
    t.integer  "subject_id",  null: false
    t.string   "type",        null: false
    t.string   "permissions"
    t.datetime "created_at"
    t.datetime "updated_at"
  end

  add_index "relations", ["observer_id", "subject_id"], name: "index_relations_on_observer_id_and_subject_id", unique: true, using: :btree
  add_index "relations", ["observer_id"], name: "index_relations_on_observer_id", using: :btree
  add_index "relations", ["subject_id"], name: "index_relations_on_subject_id", using: :btree

  create_table "themes", force: :cascade do |t|
    t.uuid     "uuid",       default: "uuid_generate_v4()", null: false
    t.string   "color"
    t.string   "wallpaper"
    t.string   "type"
    t.datetime "created_at"
    t.datetime "updated_at"
  end

  add_index "themes", ["uuid"], name: "index_themes_on_uuid", using: :btree

  create_table "users", force: :cascade do |t|
    t.uuid     "uuid",            default: "uuid_generate_v4()", null: false
    t.string   "email",                                          null: false
    t.string   "password_digest",                                null: false
    t.string   "gender",          default: "male"
    t.string   "fullname"
    t.string   "type_of_therapy"
    t.boolean  "public",          default: false,                null: false
    t.boolean  "professional",    default: false,                null: false
    t.datetime "created_at"
    t.datetime "updated_at"
    t.integer  "theme_id"
  end

  add_index "users", ["email"], name: "index_users_on_email", unique: true, using: :btree
  add_index "users", ["theme_id"], name: "index_users_on_theme_id", using: :btree
  add_index "users", ["uuid"], name: "index_users_on_uuid", using: :btree

  add_foreign_key "access_tokens", "clients", on_update: :restrict, on_delete: :restrict
  add_foreign_key "access_tokens", "users", on_update: :restrict, on_delete: :restrict
  add_foreign_key "answers", "journals", on_update: :cascade, on_delete: :cascade
  add_foreign_key "answers", "questions", on_update: :restrict, on_delete: :restrict
  add_foreign_key "devices", "users", on_update: :restrict, on_delete: :restrict
  add_foreign_key "invites", "users", on_update: :restrict, on_delete: :restrict
  add_foreign_key "journals", "clients", on_update: :restrict, on_delete: :restrict
  add_foreign_key "journals", "users", on_update: :restrict, on_delete: :restrict
  add_foreign_key "keywords", "users", on_update: :restrict, on_delete: :restrict
  add_foreign_key "notes", "users", on_update: :restrict, on_delete: :restrict
  add_foreign_key "questionaires_users", "questionaires", on_update: :restrict, on_delete: :restrict
  add_foreign_key "questionaires_users", "users", on_update: :restrict, on_delete: :restrict
  add_foreign_key "users", "themes", on_update: :restrict, on_delete: :restrict
end
