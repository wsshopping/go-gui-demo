// Package main 提供 Fyne GUI 框架的核心功能演示
// 包含窗口管理、菜单系统、主题切换等模块示例
package main

import (
    // 标准库导入
    "fmt"
    "log"
    "net/url"

    // Fyne 核心模块
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/cmd/fyne_settings/settings"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/driver/desktop"
    "fyne.io/fyne/v2/theme"
    "fyne.io/fyne/v2/widget"

    // 演示程序数据模块
    "fyne.io/demo/data"
    "fyne.io/demo/tutorials"
)

// 全局常量定义
const (
    preferenceCurrentTutorial = "currentTutorial" // 存储当前教程项的配置键
)

// 全局变量声明
var (
    topWindow fyne.Window // 主窗口引用
)

// main 应用程序入口
// 初始化 Fyne 应用，创建主窗口并设置界面布局
func main() {
    // 创建带ID的应用实例（支持配置持久化）[2](@ref)
    a := app.NewWithID("io.fyne.demo")
    a.SetIcon(data.FyneLogo) // 设置应用图标
    
    makeTray(a)  // 初始化系统托盘
    logLifecycle(a) // 注册生命周期监听

    // 创建主窗口
    w := a.NewWindow("Fyne Demo")
    topWindow = w
    w.SetMainMenu(makeMenu(a, w)) // 构建菜单系统[2](@ref)
    w.SetMaster() // 设置为主窗口（关闭时退出应用）

    // 内容区域布局
    content := container.NewStack() // 使用堆叠容器管理教程内容
    title := widget.NewLabel("Component name") // 教程标题
    intro := widget.NewLabel("An introduction would probably go\nhere, as well as a") 
    intro.Wrapping = fyne.TextWrapWord // 启用自动换行

    // 顶部区域布局（标题 + 分隔线 + 简介）
    top := container.NewVBox(title, widget.NewSeparator(), intro)

    // 教程切换函数
    setTutorial := func(t tutorials.Tutorial) {
        // 移动端处理逻辑
        if fyne.CurrentDevice().IsMobile() {
            child := a.NewWindow(t.Title)
            topWindow = child
            child.SetContent(t.View(topWindow))
            child.Show()
            child.SetOnClosed(func() { topWindow = w })
            return
        }

        // 桌面端更新界面
        title.SetText(t.Title)
        if len(t.Intro) == 0 { // Markdown内容处理
            top.Hide()
        } else {
            intro.SetText(t.Intro)
            top.Show()
        }

        content.Objects = []fyne.CanvasObject{t.View(w)}
        content.Refresh()
    }

    // 主内容布局
    tutorial := container.NewBorder(top, nil, nil, nil, content)
    if fyne.CurrentDevice().IsMobile() {
        w.SetContent(makeNav(setTutorial, false)) // 移动端导航
    } else {
        // 桌面端分栏布局（导航栏占20%宽度）[2](@ref)
        split := container.NewHSplit(makeNav(setTutorial, true), tutorial)
        split.Offset = 0.2
        w.SetContent(split)
    }

    w.Resize(fyne.NewSize(640, 460)) // 初始窗口尺寸
    w.ShowAndRun() // 显示窗口并启动事件循环
}

// logLifecycle 注册应用生命周期事件监听
func logLifecycle(a fyne.App) {
    a.Lifecycle().SetOnStarted(func() {
        log.Println("Lifecycle: Started") // 应用启动日志
    })
    a.Lifecycle().SetOnStopped(func() {
        log.Println("Lifecycle: Stopped") // 应用终止日志
    })
    a.Lifecycle().SetOnEnteredForeground(func() {
        log.Println("Lifecycle: Entered Foreground") // 进入前台
    })
    a.Lifecycle().SetOnExitedForeground(func() {
        log.Println("Lifecycle: Exited Foreground") // 退出前台
    })
}

// makeMenu 创建主菜单系统[2](@ref)
func makeMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {
    // 文件菜单项
    newItem := fyne.NewMenuItem("New", nil)
    checkedItem := fyne.NewMenuItem("Checked", nil)
    checkedItem.Checked = true // 默认勾选状态
    disabledItem := fyne.NewMenuItem("Disabled", nil)
    disabledItem.Disabled = true // 禁用状态

    // 新建子菜单
    otherItem := fyne.NewMenuItem("Other", nil)
    mailItem := fyne.NewMenuItem("Mail", func() { 
        fmt.Println("Menu New->Other->Mail") 
    })
    mailItem.Icon = theme.MailComposeIcon() // 图标设置
    otherItem.ChildMenu = fyne.NewMenu("",
        fyne.NewMenuItem("Project", func() { 
            fmt.Println("Menu New->Other->Project") 
        }),
        mailItem,
    )

    // 新建菜单项
    fileItem := fyne.NewMenuItem("File", func() { 
        fmt.Println("Menu New->File") 
    })
    fileItem.Icon = theme.FileIcon()
    dirItem := fyne.NewMenuItem("Directory", func() { 
        fmt.Println("Menu New->Directory") 
    })
    dirItem.Icon = theme.FolderIcon()
    newItem.ChildMenu = fyne.NewMenu("",
        fileItem,
        dirItem,
        otherItem,
    )

    // 设置窗口
    openSettings := func() {
        settingsWindow := a.NewWindow("Fyne Settings")
        settingsWindow.SetContent(settings.NewSettings().LoadAppearanceScreen(w))
        settingsWindow.Resize(fyne.NewSize(440, 520))
        settingsWindow.Show()
    }

    // 关于窗口
    showAbout := func() {
        aboutWindow := a.NewWindow("About")
        aboutWindow.SetContent(widget.NewLabel("About Fyne Demo app..."))
        aboutWindow.Show()
    }

    // 帮助菜单项
    aboutItem := fyne.NewMenuItem("About", showAbout)
    settingsItem := fyne.NewMenuItem("Settings", openSettings)
    settingsShortcut := &desktop.CustomShortcut{
        KeyName: fyne.KeyComma, 
        Modifier: fyne.KeyModifierShortcutDefault,
    }
    settingsItem.Shortcut = settingsShortcut
    w.Canvas().AddShortcut(settingsShortcut, func(shortcut fyne.Shortcut) {
        openSettings()
    })

    // 编辑菜单项（剪切/复制/粘贴）
    cutItem := createClipboardItem("Cut", &fyne.ShortcutCut{}, a, w)
    copyItem := createClipboardItem("Copy", &fyne.ShortcutCopy{}, a, w)
    pasteItem := createClipboardItem("Paste", &fyne.ShortcutPaste{}, a, w)

    // 查找功能
    findItem := fyne.NewMenuItem("Find", func() { 
        fmt.Println("Menu Find") 
    })
    findItem.Shortcut = &desktop.CustomShortcut{
        KeyName: fyne.KeyF,
        Modifier: fyne.KeyModifierShortcutDefault | fyne.KeyModifierAlt |
            fyne.KeyModifierShift | fyne.KeyModifierControl | fyne.KeyModifierSuper,
    }

    // 帮助菜单
    helpMenu := fyne.NewMenu("Help",
        createLinkItem("Documentation", "https://developer.fyne.io"),
        createLinkItem("Support", "https://fyne.io/support/"),
        fyne.NewMenuItemSeparator(),
        createLinkItem("Sponsor", "https://fyne.io/sponsor/"),
    )

    // 菜单组装
    file := fyne.NewMenu("File", newItem, checkedItem, disabledItem)
    if !fyne.CurrentDevice().IsMobile() {
        file.Items = append(file.Items, fyne.NewMenuItemSeparator(), settingsItem)
    }
    file.Items = append(file.Items, aboutItem)
    
    mainMenu := fyne.NewMainMenu(
        file,
        fyne.NewMenu("Edit", cutItem, copyItem, pasteItem, 
            fyne.NewMenuItemSeparator(), findItem),
        helpMenu,
    )

    // 动态菜单状态更新
    checkedItem.Action = func() {
        checkedItem.Checked = !checkedItem.Checked
        mainMenu.Refresh()
    }
    return mainMenu
}

// makeTray 创建系统托盘图标[2](@ref)
func makeTray(a fyne.App) {
    if desk, ok := a.(desktop.App); ok {
        trayItem := fyne.NewMenuItem("Hello", func() {
            log.Println("System tray menu tapped")
        })
        trayItem.Icon = theme.HomeIcon()
        menu := fyne.NewMenu("Hello World", trayItem)
        desk.SetSystemTrayMenu(menu)
    }
}

// makeNav 创建导航树[2](@ref)
func makeNav(setTutorial func(tutorial tutorials.Tutorial), loadPrevious bool) fyne.CanvasObject {
    a := fyne.CurrentApp()

    // 树形导航组件配置
    tree := &widget.Tree{
        ChildUIDs: func(uid string) []string {
            return tutorials.TutorialIndex[uid]
        },
        IsBranch: func(uid string) bool {
            children, ok := tutorials.TutorialIndex[uid]
            return ok && len(children) > 0
        },
        CreateNode: func(branch bool) fyne.CanvasObject {
            return widget.NewLabel("Collection Widgets")
        },
        UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
            if t, ok := tutorials.Tutorials[uid]; ok {
                obj.(*widget.Label).SetText(t.Title)
            }
        },
        OnSelected: func(uid string) {
            if t, ok := tutorials.Tutorials[uid]; ok {
                a.Preferences().SetString(preferenceCurrentTutorial, uid)
                setTutorial(t)
            }
        },
    }

    // 加载上次选中的教程项
    if loadPrevious {
        currentPref := a.Preferences().StringWithFallback(
            preferenceCurrentTutorial, "welcome")
        tree.Select(currentPref)
    }

    // 主题切换按钮
    themes := container.NewGridWithColumns(2,
        widget.NewButton("Dark", func() {
            a.Settings().SetTheme(theme.DarkTheme())
        }),
        widget.NewButton("Light", func() {
            a.Settings().SetTheme(theme.LightTheme())
        }),
    )

    return container.NewBorder(nil, themes, nil, nil, tree)
}

// 辅助函数 ---------------------------------------------------

// createClipboardItem 创建剪贴板操作菜单项
func createClipboardItem(label string, shortcut fyne.Shortcut, 
    a fyne.App, w fyne.Window) *fyne.MenuItem {
    item := fyne.NewMenuItem(label, func() {
        shortcutFocused(shortcut, a.Clipboard(), w.Canvas().Focused())
    })
    item.Shortcut = shortcut
    return item
}

// createLinkItem 创建带URL链接的菜单项
func createLinkItem(label, urlStr string) *fyne.MenuItem {
    return fyne.NewMenuItem(label, func() {
        u, _ := url.Parse(urlStr)
        _ = fyne.CurrentApp().OpenURL(u)
    })
}

// shortcutFocused 处理快捷键事件[2](@ref)
func shortcutFocused(s fyne.Shortcut, cb fyne.Clipboard, f fyne.Focusable) {
    // 设置剪贴板引用
    switch sh := s.(type) {
    case *fyne.ShortcutCopy: sh.Clipboard = cb
    case *fyne.ShortcutCut: sh.Clipboard = cb
    case *fyne.ShortcutPaste: sh.Clipboard = cb
    }
    
    // 触发焦点组件的快捷键处理
    if focused, ok := f.(fyne.Shortcutable); ok {
        focused.TypedShortcut(s)
    }
}
