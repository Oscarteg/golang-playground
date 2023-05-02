package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"gorm.io/driver/sqlite" // Sqlite driver based on GGO

	"gorm.io/gorm"

	"github.com/go-chi/chi/v5/middleware"
)

type Article struct {
	gorm.Model `json:"model"`
	Title      string `gorm:"serializer:json"`
}

// Create an exported global variable to hold the database connection pool.
var db *gorm.DB

func main() {
	r := chi.NewRouter()

	var err error

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	db, err = gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Article{})

	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(DatabaseContext)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	r.Route("/articles", func(r chi.Router) {
		r.With(paginate).Get("/", listArticles)
		r.Post("/", createArticle)

		r.Route("/{articleID}", func(r chi.Router) {
			r.Use(ArticleContext)
			r.Get("/", getArticle)
		})
	})

	fmt.Printf("Server is running on port 3333")

	http.ListenAndServe(":3333", r)
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	db := ctx.Value("db").(*gorm.DB)
	article := Article{
		Title: "test",
	}

	if result := db.Create(&article); result.Error != nil {
		http.Error(w, http.StatusText(422), 422)
		return
	}
	w.Write([]byte(fmt.Sprintf("%s", "asd")))
}

func listArticles(w http.ResponseWriter, r *http.Request) {
	var articles []Article
	db := r.Context().Value("db").(*gorm.DB)

	if err := db.Find(&articles); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Write([]byte(fmt.Sprintf("articles:%s", articles)))
}

func ArticleContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		articleId := chi.URLParam(r, "articleId")
		var article Article
		if err := db.First(&article, articleId); err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		ctx = context.WithValue(ctx, "article", article)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func DatabaseContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "DB", &db)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	article, ok := ctx.Value("article").(*Article)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}
	w.Write([]byte(fmt.Sprintf("title:%s", article.Title)))

}

func getArticle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	article, ok := ctx.Value("article").(*Article)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}
	w.Write([]byte(fmt.Sprintf("title:%s", article.Title)))
}

// paginate is a stub, but very possible to implement middleware logic
// to handle the request params for handling a paginated request.
func paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// just a stub.. some ideas are to look at URL query params for something like
		// the page number, or the limit, and send a query cursor down the chain
		next.ServeHTTP(w, r)
	})
}

type ArticleRepository struct {
	DB *gorm.DB
}

func (r *ArticleRepository) FindAll() ([]Article, error) {
	var articles []Article
	r.DB.Find(&articles)
	if err := r.DB.Find(&articles).Error; err != nil {
		return nil, err
	}

	return articles, nil
}

type ArticleService struct {
	ArticleRepository *ArticleRepository
}
