package lokasiasetmodel

import (
	"database/sql"
	"fmt"
	"inventaris/config"
	"inventaris/entities"

	"github.com/google/uuid"
)

func GetAll() []entities.LokasiAset {
	rows, err := config.DB.Query(
		`SELECT 
			la.*, 
			at.nama_aset, 
			at.kode_aset,
			at.path,
			bb.nama_bidang,
			pp.name,
			pp.nip,
			pp.foto
		FROM 
			lokasi_aset la
		JOIN 
			aset_tik at ON la.aset_id = at.id
		JOIN 
			bidang bb ON la.bidang_id = bb.id
		JOIN 
			pegawai pp ON la.pegawai_id = pp.id
		GROUP BY
			la.pegawai_id
		ORDER BY 
			la.updated_at DESC;
		`)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var lokasiAsets []entities.LokasiAset

	for rows.Next() {
		var lokasiaset entities.LokasiAset

		if err := rows.Scan(
			&lokasiaset.Id,
			&lokasiaset.Aset_id,
			&lokasiaset.Bidang_id,
			&lokasiaset.Pegawai_id,
			&lokasiaset.Tanggal_Perolehan,
			&lokasiaset.Tanggal_Selesai,
			&lokasiaset.Jenis_Pemanfaatan,
			&lokasiaset.Keterangan,
			&lokasiaset.Created_At,
			&lokasiaset.Updated_At,
			&lokasiaset.Nama_Aset,
			&lokasiaset.Kode_Aset,
			&lokasiaset.Path,
			&lokasiaset.Nama_Bidang,
			&lokasiaset.Nama_Pegawai,
			&lokasiaset.Nip_Pegawai,
			&lokasiaset.Foto_Pegawai,
		); err != nil {
			panic(err.Error())
		}

		lokasiAsets = append(lokasiAsets, lokasiaset)
	}

	return lokasiAsets
}

func Create(lokasiaset entities.LokasiAset) (bool, error) {
	newUUID := uuid.New()
	result, err := config.DB.Exec(
		`INSERT INTO
		lokasi_aset (
		id, 
		aset_id, 
		bidang_id, 
		pegawai_id, 
		tanggal_perolehan, 
		tanggal_selesai, 
		jenis_pemanfaatan, 
		keterangan, 
		created_at, 
		updated_at
		) VALUES (
		 ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		newUUID,
		lokasiaset.Aset_id,
		lokasiaset.Bidang_id,
		lokasiaset.Pegawai_id,
		lokasiaset.Tanggal_Perolehan,
		lokasiaset.Tanggal_Selesai,
		lokasiaset.Jenis_Pemanfaatan,
		lokasiaset.Keterangan,
		lokasiaset.Created_At,
		lokasiaset.Updated_At,
	)

	if err != nil {
		panic(err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	return rowsAffected > 0, nil
}

func Detail(id string) (entities.LokasiAset, error) {
	row := config.DB.QueryRow(`
	SELECT 
			la.*, 
			at.nama_aset, 
			at.kode_aset,
			at.path,
			bb.nama_bidang,
			pp.name,
			pp.nip, 
			pp.foto,
			pp.jenis_pegawai,
			pp.jabatan
		FROM 
			lokasi_aset la
		JOIN 
			aset_tik at ON la.aset_id = at.id
		JOIN 
			bidang bb ON la.bidang_id = bb.id
		JOIN 
			pegawai pp ON la.pegawai_id = pp.id
		WHERE 
			la.id = ?
		ORDER BY 
			la.updated_at DESC`, id)

	var lokasiaset entities.LokasiAset

	if err := row.Scan(
		&lokasiaset.Id,
		&lokasiaset.Aset_id,
		&lokasiaset.Bidang_id,
		&lokasiaset.Pegawai_id,
		&lokasiaset.Tanggal_Perolehan,
		&lokasiaset.Tanggal_Selesai,
		&lokasiaset.Jenis_Pemanfaatan,
		&lokasiaset.Keterangan,
		&lokasiaset.Created_At,
		&lokasiaset.Updated_At,
		&lokasiaset.Nama_Aset,
		&lokasiaset.Kode_Aset,
		&lokasiaset.Path,
		&lokasiaset.Nama_Bidang,
		&lokasiaset.Nama_Pegawai,
		&lokasiaset.Nip_Pegawai,
		&lokasiaset.Foto_Pegawai,
		&lokasiaset.Jenis_Pegawai,
		&lokasiaset.Jabatan,
	); err != nil {
		if err == sql.ErrNoRows {
			return lokasiaset, fmt.Errorf("lokasi aset tidak ditemukan dengan id %s", id)
		}
		return lokasiaset, fmt.Errorf("failed to retrieve lokasi aset: %w", err)
	}

	return lokasiaset, nil
}

func DaftarAset(pegawai_id string) []entities.LokasiAset {
	rows, err := config.DB.Query(`
	SELECT 
			la.*, 
			at.nama_aset, 
			at.kode_aset,
			at.path,
			bb.nama_bidang,
			pp.nip, 
			pp.name
		FROM 
			lokasi_aset la
		JOIN 
			aset_tik at ON la.aset_id = at.id
		JOIN 
			bidang bb ON la.bidang_id = bb.id
		JOIN 
			pegawai pp ON la.pegawai_id = pp.id
		WHERE 
			la.pegawai_id = ?
		ORDER BY 
			la.updated_at DESC`, pegawai_id)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var lokasiasets []entities.LokasiAset

	for rows.Next() {
		var lokasiaset entities.LokasiAset

		if err := rows.Scan(
			&lokasiaset.Id,
			&lokasiaset.Aset_id,
			&lokasiaset.Bidang_id,
			&lokasiaset.Pegawai_id,
			&lokasiaset.Tanggal_Perolehan,
			&lokasiaset.Tanggal_Selesai,
			&lokasiaset.Jenis_Pemanfaatan,
			&lokasiaset.Keterangan,
			&lokasiaset.Created_At,
			&lokasiaset.Updated_At,
			&lokasiaset.Nama_Aset,
			&lokasiaset.Kode_Aset,
			&lokasiaset.Path,
			&lokasiaset.Nama_Bidang,
			&lokasiaset.Nama_Pegawai,
			&lokasiaset.Nip_Pegawai,
		); err != nil {
			panic(err.Error())
		}

		lokasiasets = append(lokasiasets, lokasiaset)
	}

	return lokasiasets
}

func Update(id string, lokasiaset entities.LokasiAset) (bool, error) {
	query, err := config.DB.Exec(`
	UPDATE lokasi_aset
		SET
			aset_id = ?,
			bidang_id = ?,
			pegawai_id = ?,
			tanggal_perolehan = ?,
			tanggal_selesai = ?,
			jenis_pemanfaatan = ?,
			keterangan = ?,
			updated_at = ?
		WHERE id = ?
	`,
		lokasiaset.Aset_id,
		lokasiaset.Bidang_id,
		lokasiaset.Pegawai_id,
		lokasiaset.Tanggal_Perolehan,
		lokasiaset.Tanggal_Selesai,
		lokasiaset.Jenis_Pemanfaatan,
		lokasiaset.Keterangan,
		lokasiaset.Updated_At,
		id,
	)
	if err != nil {
		panic(err.Error())
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	return result > 0, nil
}

func Delete(id string) error {
	_, err := config.DB.Exec(`DELETE FROM lokasi_aset WHERE id = ?`, id)
	return err
}
