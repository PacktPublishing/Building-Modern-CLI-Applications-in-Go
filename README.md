# Building Modern CLI Applications in Go

<a href="https://www.packtpub.com/product/building-modern-cli-applications-in-go/9781804611654?utm_source=github&utm_medium=repository&utm_campaign=9781804611654"><img src="https://static.packt-cdn.com/products/9781804611654/cover/smaller" alt="Building Modern CLI Applications in Go" height="256px" align="right"></a>

This is the code repository for [Building Modern CLI Applications in Go](https://www.packtpub.com/product/building-modern-cli-applications-in-go/9781804611654?utm_source=github&utm_medium=repository&utm_campaign=9781804611654), published by Packt.

**Develop next-level CLIs to improve user experience, increase platform usage, and maximize production**

## What is this book about?
Although graphical user interfaces (GUIs) are intuitive and user-friendly, nothing beats a command-line

This book covers the following exciting features:
* Master the Go code structure, testing, and other essentials
* Add a colorful dashboard to your CLI using engaging ASCII banners
* Use Cobra, Viper, and other frameworks to give your CLI an edge
* Handle inputs, API commands, errors, and timeouts like a pro
* Target builds for specific platforms the right way using build tags
* Build with empathy, using easy bug submission and traceback
* Containerize, distribute, and publish your CLIs quickly and easily

If you feel this book is for you, get your [copy](https://www.amazon.com/dp/1804611654) today!

<a href="https://www.packtpub.com/?utm_source=github&utm_medium=banner&utm_campaign=GitHubBanner"><img src="https://raw.githubusercontent.com/PacktPublishing/GitHub/master/GitHub.png" 
alt="https://www.packtpub.com/" border="5" /></a>

## Instructions and Navigations
All of the code is organized into folders. For example, Chapter04.

The code will look like the following:
```
func init() {
    audioCmd.Flags().StringP("filename", "f", "", "audiofile")
    uploadCmd.AddCommand(audioCmd)
}
```

**Following is what you need for this book:**
This book is for beginner- and intermediate-level Golang developers who take an interest in developing CLIs and enjoy learning by doing. You'll need an understanding of basic Golang programming concepts, but will require no prior knowledge of CLI design and development. This book helps you join a community of CLI developers and distribute within the popular Homebrew package management tool.

With the following software and hardware list you can run all code files present in the book (Chapter 1-14).
### Software and Hardware List
| Chapter | Software required | OS required |
| -------- | ------------------------------------ | ----------------------------------- |
| 1-14 | Go 1.19 | Windows, Mac OS X, and Linux (Any) |
| 1-14 | Cobra CLI | Windows, Mac OS X, and Linux (Any) |
| 1-14 | Docker | Windows, Mac OS X, and Linux (Any) |
| 1-14 | Docker Compose | Windows, Mac OS X, and Linux (Any) |
| 1-14 | GoReleaser CLI | Windows, Mac OS X, and Linux (Any) |

We also provide a PDF file that has color images of the screenshots/diagrams used in this book. [Click here to download it](https://packt.link/F4Fus).

### Related products
* Event-Driven Architecture in Golang [[Packt]](https://www.packtpub.com/product/event-driven-architecture-in-golang/9781803238012?utm_source=github&utm_medium=repository&utm_campaign=9781803238012) [[Amazon]](https://www.amazon.com/dp/1803238011)

* Domain-Driven Design with Golang [[Packt]](https://www.packtpub.com/product/domain-driven-design-with-golang/9781804613450?utm_source=github&utm_medium=repository&utm_campaign=9781804613450) [[Amazon]](https://www.amazon.com/dp/1804613452)

## Get to know the Author
**Marian Montagnino**
is a Senior Software Engineer at Netflix with over 20 years of experience. Since the early nineties, when her family first got a home computer, she has been using the terminal and command line applications to navigate through text-based systems. In 1995, she held her first job as a SysOp, or system operator, for Real of Mirage BBS in Fair Lawn, NJ. Her early years discovering technology inspired her to continue learning about computers. She received her Dual Computer Science and Mathematics of Operations Research BSc from Rensselaer Polytechnic Institute and her Applied Mathematics MSc from Stevens Institute of Technology.

### Download a free PDF

 <i>If you have already purchased a print or Kindle version of this book, you can get a DRM-free PDF version at no cost.<br>Simply click on the link to claim your free PDF.</i>
<p align="center"> <a href="https://packt.link/free-ebook/9781804611654">https://packt.link/free-ebook/9781804611654 </a> </p>