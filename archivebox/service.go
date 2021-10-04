package archivebox

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/input"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
)

type Config struct {
	Path     string
	Username string
	Password string
}

type API interface {
	ArchiveLink(url string) (string, string, error)
}

func New(cfg Config) API {
	return &service{cfg: cfg}
}

type service struct {
	cfg Config
}

func (s *service) ArchiveLink(url string) (string, string, error) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("ignore-certificate-errors", "1"),
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := context.WithTimeout(allocCtx, time.Second*30)

	ctx, cancel = chromedp.NewContext(ctx, chromedp.WithLogf(log.Printf))
	defer cancel()

	var title, fileUrl string
	if err := chromedp.Run(ctx,
		chromedp.Navigate(fmt.Sprintf("%s/admin/login/", s.cfg.Path)),
		chromedp.WaitVisible("#id_username", chromedp.ByQuery),
		chromedp.SendKeys("#id_username", s.cfg.Username, chromedp.ByQuery),
		chromedp.SendKeys("#id_password", s.cfg.Password, chromedp.ByQuery),
		chromedp.Submit("#id_username", chromedp.ByQuery),
		chromedp.WaitVisible("#recent-actions-module", chromedp.ByQuery),
		chromedp.Navigate(fmt.Sprintf("%s/add/", s.cfg.Path)),
		chromedp.WaitVisible("#id_url", chromedp.ByQuery),
		chromedp.SendKeys("#id_url", url, chromedp.ByQuery),
		chromedp.Click("#id_depth_0", chromedp.ByQuery),
		chromedp.QueryAfter("#id_archive_methods > option[value=\"title\"]", func(ctx context.Context, id runtime.ExecutionContextID, nodes ...*cdp.Node) error {
			return chromedp.MouseClickNode(nodes[0], chromedp.ButtonModifiers(input.ModifierCtrl)).Do(ctx)
		}, chromedp.ByQuery),
		chromedp.QueryAfter("#id_archive_methods > option[value=\"singlefile\"]", func(ctx context.Context, id runtime.ExecutionContextID, nodes ...*cdp.Node) error {
			return chromedp.MouseClickNode(nodes[0], chromedp.ButtonModifiers(input.ModifierCtrl)).Do(ctx)
		}, chromedp.ByQuery),
		chromedp.Click("#submit", chromedp.ByQuery),
		chromedp.WaitVisible("pre#stdout", chromedp.ByQuery),
		chromedp.Navigate(fmt.Sprintf("%s/public/", s.cfg.Path)),
		chromedp.WaitVisible("#searchbar", chromedp.ByQuery),
		chromedp.SendKeys("#searchbar", url, chromedp.ByQuery),
		chromedp.SendKeys("#searchbar", kb.Enter, chromedp.ByQuery),
		chromedp.WaitVisible("#searchbar", chromedp.ByQuery),
		chromedp.Text(".title-col > a:nth-child(2) > span:nth-child(1)", &title, chromedp.ByQuery),
		chromedp.AttributeValue("a[title='singlefile']", "href", &fileUrl, nil, chromedp.ByQuery),
	); err != nil {
		return "", "", err
	}

	return title, fmt.Sprintf("%s%s", s.cfg.Path, fileUrl), nil
}
