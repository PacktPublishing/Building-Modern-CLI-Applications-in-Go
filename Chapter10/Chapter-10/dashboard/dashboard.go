package dashboard

import (
	"context"
	"fmt"
	"time"

	"github.com/mum4k/termdash"

	"github.com/mum4k/termdash/align"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"

	"github.com/mum4k/termdash/container/grid"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/mouse"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgets/button"
	"github.com/mum4k/termdash/widgets/text"
)

func Basic() error {
	terminalLayer, err := tcell.New()
	if err != nil {
		return err
	}
	defer terminalLayer.Close()

	containerLayer, err := container.New(terminalLayer)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := termdash.Run(ctx, terminalLayer, containerLayer); err != nil {
		return err
	}
	return nil
}

func BasicWithRedraw() error {
	terminalLayer, err := tcell.New()
	if err != nil {
		return err
	}
	defer terminalLayer.Close()

	containerLayer, err := container.New(terminalLayer)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()
	if err := termdash.Run(ctx, terminalLayer, containerLayer, termdash.RedrawInterval(5*time.Second)); err != nil {
		return err
	}
	return nil
}

func Redraw() error {
	terminalLayer, err := tcell.New()
	if err != nil {
		return err
	}
	defer terminalLayer.Close()

	containerLayer, err := container.New(terminalLayer)
	if err != nil {
		return err
	}

	termController, err := termdash.NewController(terminalLayer, containerLayer)
	if err != nil {
		return err
	}
	defer termController.Close()

	if err := termController.Redraw(); err != nil {
		return fmt.Errorf("error redrawing dashboard: %v", err)
	}
	return nil
}

func ErrorHandlingOption() error {
	terminalLayer, err := tcell.New()
	if err != nil {
		return err
	}
	defer terminalLayer.Close()

	containerLayer, err := container.New(terminalLayer)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	errHandler := func(err error) {
		fmt.Printf("runtime error: %v", err)
	}
	defer cancel()
	if err := termdash.Run(ctx, terminalLayer, containerLayer, termdash.ErrorHandler(errHandler)); err != nil {
		return err
	}
	return nil
}

func KeyboardSubscriber() error {
	terminalLayer, err := tcell.New()
	if err != nil {
		return err
	}
	defer terminalLayer.Close()

	containerLayer, err := container.New(terminalLayer)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)

	keyboardSubscriber := func(k *terminalapi.Keyboard) {
		switch k.Key {
		case 'q':
		case 'Q':
			cancel()
		}
	}
	if err := termdash.Run(ctx, terminalLayer, containerLayer, termdash.KeyboardSubscriber(keyboardSubscriber)); err != nil {
		return fmt.Errorf("error running termdash with keyboard subscriber: %v", err)
	}
	return nil
}

func MouseSubscriber() error {
	terminalLayer, err := tcell.New()
	if err != nil {
		return err
	}
	defer terminalLayer.Close()

	containerLayer, err := container.New(terminalLayer)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)

	mouseClick := func(m *terminalapi.Mouse) {
		switch m.Button {
		case mouse.ButtonRight:
			cancel()
		case mouse.ButtonLeft:
			// when the left mouse button is clicked
		case mouse.ButtonMiddle:
			// when the middle mouse button is clicked
		}
	}

	if err := termdash.Run(ctx, terminalLayer, containerLayer, termdash.MouseSubscriber(mouseClick)); err != nil {
		return fmt.Errorf("error running termdash with mouse subscriber: %v", err)
	}
	return nil
}

func BinaryTree() error {
	terminalLayer, err := tcell.New(tcell.ColorMode(terminalapi.ColorMode256),
		tcell.ClearStyle(cell.ColorYellow, cell.ColorNavy))
	if err != nil {
		return err
	}
	defer terminalLayer.Close()
	leftContainer := container.Left(
		container.Border(linestyle.Light),
	)
	rightContainer :=
		container.Right(container.SplitHorizontal(
			container.Top(
				container.Border(linestyle.Light),
			),
			container.Bottom(
				container.SplitVertical(
					container.Left(
						container.Border(linestyle.Light),
					),
					container.Right(
						container.Border(linestyle.Light),
					),
				),
			),
		))
	containerLayer, err := container.New(
		terminalLayer,
		container.SplitVertical(
			leftContainer,
			rightContainer,
			container.SplitPercent(60),
		),
	)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	keyboardSubscriber := func(k *terminalapi.Keyboard) {
		switch k.Key {
		case 'q':
		case 'Q':
			cancel()
		}
	}

	if err := termdash.Run(ctx, terminalLayer, containerLayer, termdash.KeyboardSubscriber(keyboardSubscriber)); err != nil {
		return fmt.Errorf("error running termdash with mouse subscriber: %v", err)
	}
	return nil
}

func Grid() error {
	t, err := tcell.New()
	if err != nil {
		return fmt.Errorf("error creating tcell: %v", err)
	}
	rollingText, err := text.New(text.RollContent())
	if err != nil {
		return fmt.Errorf("error creating rolling text: %v", err)
	}
	err = rollingText.Write("...")
	if err != nil {
		return fmt.Errorf("error writing text: %v", err)
	}
	builder := grid.New()
	builder.Add(
		grid.ColWidthPerc(60,
			grid.Widget(rollingText,
				container.Border(linestyle.Light),
			),
		),
	)
	builder.Add(
		grid.RowHeightPerc(50,
			grid.Widget(rollingText,
				container.Border(linestyle.Light),
			),
		),
	)
	builder.Add(
		grid.ColWidthPerc(20,
			grid.Widget(rollingText,
				container.Border(linestyle.Light),
			),
		),
	)
	builder.Add(
		grid.ColWidthPerc(20,
			grid.Widget(rollingText,
				container.Border(linestyle.Light),
			),
		),
	)
	gridOpts, err := builder.Build()
	if err != nil {
		return fmt.Errorf("error creating builder: %v", err)
	}

	c, err := container.New(t, gridOpts...)
	if err != nil {
		return fmt.Errorf("error creating container: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	if err := termdash.Run(ctx, t, c); err != nil {
		return fmt.Errorf("error running termdash: %v", err)
	}
	return nil
}

func DynamicLayout() error {
	t, err := tcell.New()
	if err != nil {
		return fmt.Errorf("error creating tcell: %v", err)
	}
	defer t.Close()

	b1, err := button.New("button1", func() error {
		return nil
	})
	if err != nil {
		return fmt.Errorf("error creating button: %v", err)
	}

	b2, err := button.New("button2", func() error {
		return nil
	})
	if err != nil {
		return fmt.Errorf("error creating button: %v", err)
	}

	c, err := container.New(
		t,
		container.PlaceWidget(b1),
		container.ID("123"),
	)
	if err != nil {
		return fmt.Errorf("error creating container: %v", err)
	}
	update := func(k *terminalapi.Keyboard) {
		if k.Key == 'u' || k.Key == 'U' {
			c.Update(
				"123",
				container.SplitVertical(
					container.Left(
						container.PlaceWidget(b1),
					),
					container.Right(
						container.PlaceWidget(b2),
					),
				),
			)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := termdash.Run(ctx, t, c, termdash.KeyboardSubscriber(update)); err != nil {
		return fmt.Errorf("error running termdash: %v", err)
	}
	return nil
}

func BinaryTreeWithStyle() error {
	terminalLayer, err := tcell.New(tcell.ColorMode(terminalapi.ColorMode256),
		tcell.ClearStyle(cell.ColorYellow, cell.ColorNavy))
	if err != nil {
		return err
	}
	defer terminalLayer.Close()
	b, err := button.New("click me", func() error {
		return nil
	})
	if err != nil {
		return err
	}
	leftContainer :=
		container.Left(
			container.Border(linestyle.Light),
			container.PlaceWidget(b),
			container.AlignHorizontal(align.HorizontalLeft),
		)
	rightContainer :=
		container.Right(
			container.SplitHorizontal(
				container.Top(
					container.Border(linestyle.Light),
					container.PlaceWidget(b),
					container.AlignVertical(align.VerticalTop),
				),
				container.Bottom(
					container.SplitVertical(
						container.Left(
							container.Border(linestyle.Light),
							container.PlaceWidget(b),
							container.PaddingTop(3),
							container.PaddingBottom(3),
							container.PaddingRight(3),
							container.PaddingLeft(3),
						),
						container.Right(
							container.Border(linestyle.Light),
							container.PlaceWidget(b),
							container.MarginTop(3),
							container.MarginBottom(3),
							container.MarginRight(3),
							container.MarginLeft(3),
						),
					),
				),
			),
		)
	containerLayer, err := container.New(
		terminalLayer,
		container.SplitVertical(
			leftContainer,
			rightContainer,
			container.SplitPercent(60),
		),
	)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	keyboardSubscriber := func(k *terminalapi.Keyboard) {
		switch k.Key {
		case 'q':
		case 'Q':
			cancel()
		}
	}

	if err := termdash.Run(ctx, terminalLayer, containerLayer, termdash.KeyboardSubscriber(keyboardSubscriber)); err != nil {
		return fmt.Errorf("error running termdash with mouse subscriber: %v", err)
	}
	return nil
}
