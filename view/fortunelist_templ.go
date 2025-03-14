// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package view

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"main/model"
	"main/view/layout"
)

func FortuneList(memories []model.Fortune, wishes []model.Fortune) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div><div class=\"button-container\"><a href=\"/fortunes/form?kind=memory\" class=\"btn\">I remember...</a> <a href=\"/fortunes/form?kind=wish\" class=\"btn\">I wish...</a></div><div class=\"bouncing-container\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for _, fortune := range memories {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a href=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var3 templ.SafeURL = templ.SafeURL("/fortunes/" + fortune.GetStrId())
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var3)))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"bouncing-link memory-link\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var4 string
				templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(fortune.GetTitle())
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/fortunelist.templ`, Line: 17, Col: 133}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a> ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			for _, fortune := range wishes {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a href=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var5 templ.SafeURL = templ.SafeURL("/fortunes/" + fortune.GetStrId())
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var5)))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"bouncing-link wish-link\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var6 string
				templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(fortune.GetTitle())
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/fortunelist.templ`, Line: 20, Col: 131}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><script>\n            const elements = [];\n            // Get existing bouncing elements\n            const existingElements = document.querySelectorAll('.bouncing-link');\n            // Add a new bouncing element\n            existingElements.forEach(element => {\n                const rect = element.getBoundingClientRect();\n                // Random position\n                const x = Math.random() * (window.innerWidth - rect.width);\n                const y = Math.random() * (window.innerHeight - rect.height);\n                element.style.left = `${x}px`;\n                element.style.top = `${y}px`;\n                \n                // Random speed (1-5 pixels per frame)\n                const speed = (0.5 + Math.random() * 0.5) * (Math.random() > 0.3 ? 1 : -1);\n                const speedX = speed;\n                const speedY = speed;\n\n                elements.push({\n                    element,\n                    x,\n                    y,\n                    speedX,\n                    speedY,\n                    width: rect.width,\n                    height: rect.height\n                });\n            });\n        \n            // Animation loop\n            function animate() {\n                elements.forEach(item => {\n                    // Update position\n                    item.x += item.speedX;\n                    item.y += item.speedY;\n                    \n                    // Check for collisions with walls\n                    if (item.x <= 0 || item.x + item.width >= window.innerWidth) {\n                        item.speedX = -item.speedX;\n                    }\n                    \n                    if (item.y <= 0 || item.y + item.height >= window.innerHeight) {\n                        item.speedY = -item.speedY;\n                    }\n                    \n                    // Apply new position\n                    item.element.style.left = `${item.x}px`;\n                    item.element.style.top = `${item.y}px`;\n                });\n                \n                requestAnimationFrame(animate);\n            }\n            \n            // Start animation\n            animate();\n            \n            // Handle window resizing\n            window.addEventListener('resize', () => {\n                elements.forEach(item => {\n                    // Keep elements in bounds after resize\n                    if (item.x + item.width > window.innerWidth) {\n                        item.x = window.innerWidth - item.width;\n                    }\n                    if (item.y + item.height > window.innerHeight) {\n                        item.y = window.innerHeight - item.height;\n                    }\n                });\n            });\n        </script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layout.Base().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
