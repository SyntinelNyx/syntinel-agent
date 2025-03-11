package kopia

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/kopia/kopia/fs"
	"github.com/kopia/kopia/fs/localfs"
	"github.com/kopia/kopia/repo"
	_ "github.com/kopia/kopia/repo/blob/s3"
	"github.com/kopia/kopia/snapshot"
	// "github.com/kopia/kopia/snapshot/policy"
	"github.com/kopia/kopia/snapshot/snapshotfs"
	// "github.com/kopia/kopia/snapshot/restore"
)

func CreateSnapshot(
	ctx context.Context,
	fsentry fs.Entry,
	rep repo.Repository,
	repwriter repo.RepositoryWriter,
	uploader *snapshotfs.Uploader,
	srcinfo snapshot.SourceInfo) (err error) {
	// _, err = policy.TreeForSource(ctx, repwriter, srcinfo)
	// if err != nil {
	// 	return fmt.Errorf("failed to get policy tree from source: %v", err)
	// }


	manifest, err := uploader.Upload(ctx, fsentry, nil, srcinfo)
	if err != nil {
		return fmt.Errorf("failed to upload entry: %v", err)
	}

	manid, err := snapshot.SaveSnapshot(ctx, repwriter, manifest)
	if err != nil {
		return fmt.Errorf("failed to create snapshot: %v", err)
	}
	fmt.Printf("Snapshot created successfully with ID: %s\n", manid)
	fmt.Printf("Manifest: %v\n", manifest)

	// if _, err := policy.ApplyRetentionPolicy(ctx, repwriter, srcinfo, true); err != nil {
	// 	return fmt.Errorf("failed to apply retention policy: %v", err)
	// }

	// snapshots, err := snapshot.ListSnapshots(ctx, rep, srcinfo)
	// if err != nil {
	// 	log.Fatalf("failed to list snapshots: %v\n", err)
	// }

	// snapshots, err := snapshot.LoadSnapshot(ctx, rep,)
	// if err != nil {
	// 	log.Fatalf("failed to list snapshots: %v\n", err)
	// }

	// fmt.Printf("Snapshots: %v\n", snapshots)


	if err := repwriter.Close(ctx); err != nil {
		log.Fatalf("failed to close writer: %v\n", err)
	}

	return nil
}

func OpenRepository() {
	ctx := context.Background()
	password := os.Getenv("KOPIA_REPO_PASSWORD")

	rep, err := repo.Open(ctx, "./config.json", password, nil)
	if err != nil {
		log.Fatalf("failed to open repository: %v\n", err)
	}
	defer rep.Close(ctx)



	err = repo.WriteSession(ctx, rep, repo.WriteSessionOptions{
		Purpose: "CreateSnapshot",
	},  func(ctx context.Context, w repo.RepositoryWriter) error {
		srcdir, err := filepath.Abs("./data/test")
		if err != nil {
			log.Fatalf("failed to resolve file path: %v\n", err)
		}
		srcinfo := snapshot.SourceInfo{
			Host:     rep.ClientOptions().Hostname,
			UserName: rep.ClientOptions().Username,
			Path:     filepath.Clean(srcdir),
		}
	
		entry, err := localfs.NewEntry(srcinfo.Path)
		if err != nil {
			log.Fatalf("failed to create new entry: %v\n", err)
		}
	
		uploader := snapshotfs.NewUploader(w)
	
		if err := CreateSnapshot(ctx, entry, rep, w, uploader, srcinfo); err != nil {
			log.Fatalf("failed to snapshot source: %v\n", err)
		}
	
		if err := w.Flush(ctx); err != nil {
			log.Fatalf("failed to flush writer: %v\n", err)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("failed to write session: %v\n", err)
	}

	// ctx, repwriter, err := rep.NewWriter(ctx, repo.WriteSessionOptions{
	// 	Purpose: "CreateSnapshot",
	// })
	// if err != nil {
	// 	log.Fatalf("failed to create writer: %v\n", err)
	// }

	// srcdir, err := filepath.Abs("./data/test")
	// if err != nil {
	// 	log.Fatalf("failed to resolve file path: %v\n", err)
	// }
	// srcinfo := snapshot.SourceInfo{
	// 	Host:     rep.ClientOptions().Hostname,
	// 	UserName: rep.ClientOptions().Username,
	// 	Path:     filepath.Clean(srcdir),
	// }

	// entry, err := localfs.NewEntry(srcinfo.Path)
	// if err != nil {
	// 	log.Fatalf("failed to create new entry: %v\n", err)
	// }

	// uploader := snapshotfs.NewUploader(repwriter)

	// if err := CreateSnapshot(ctx, entry, rep, repwriter, uploader, srcinfo); err != nil {
	// 	log.Fatalf("failed to snapshot source: %v\n", err)
	// }

	// if err := repwriter.Flush(ctx); err != nil {
	// 	log.Fatalf("failed to flush writer: %v\n", err)
	// }

	// snapshots, err := snapshot.LoadSnapshot(ctx, rep, 0)
	// if err != nil {
	// 	log.Fatalf("failed to load snapshots: %v\n", err)
	// }

	// fmt.Printf("Snapshots: %v\n", snapshots)

	// sources, err := snapshot.ListSources(ctx, rep)
	// if err != nil {
	// 	log.Fatalf("failed to list snapshots: %v\n", err)
	// }

	// fmt.Printf("Snapshots: %v\n", sources)

	// snapshots, err := snapshot.ListSnapshots(ctx, rep, srcinfo, )
	// if err != nil {
	// 	log.Fatalf("failed to list snapshots: %v\n", err)
	// }

	// fmt.Printf("Snapshots: %v\n", snapshots)

	// restoredir := "./data/restored"
	// if err := os.MkdirAll(restoredir, 0755); err != nil {
	// 	log.Fatalf("failed to create restore directory: %v", err)
	// }
	//
	// outputDir, err := localfs.Directory(restoredir)
	// if err != nil {
	// 	log.Fatalf("failed to open restore directory: %v", err)
	// }
	//
	// if _, err := restore.Entry(ctx, rep, outputDir, man.RootEntry, restore.Options{}); err != nil {
	// 	log.Fatalf("failed to restore snapshot: %v", err)
	// }
	//
	// fmt.Println("Snapshot restored successfully.")

}
