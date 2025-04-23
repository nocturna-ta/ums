ALTER TABLE "kpu_provinsi" add column "telephone" VARCHAR(20) DEFAULT '';
ALTER TABLE "kpu_kota" add column "telephone" VARCHAR(20) DEFAULT '';
ALTER TABLE "kpu_provinsi" add column "registered_at" TIMESTAMP(6) WITH TIME ZONE DEFAULT now();
ALTER TABLE "kpu_kota" add column "registered_at" TIMESTAMP(6) WITH TIME ZONE DEFAULT now();