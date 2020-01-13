package templating

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"testing"
)

// Tests are no where near complete. Just used it for basic smoke testing
// TODO needs more tests

func TestNewTmpl(t *testing.T) {
	type args struct {
		dir   string
		ext   string
		layoutDir string
		devel bool
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		templateNames []string
	}{
		{
			name: "Init Develop",
			args: args{dir: "../../templates", layoutDir: "../../templates/layouts", ext: ".gohtml", devel: true},
			wantErr: false,
			templateNames: []string{"accounts/index", "users/show", "index"},

		},
		{
			name: "Init Not Develop",
			args: args{dir: "../../templates", layoutDir: "../../templates/layouts", ext: ".gohtml", devel: false},
			wantErr: false,
			templateNames: []string{"accounts/index", "users/show", "index"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTmpl, err := NewTemplateManager(tt.args.dir, tt.args.layoutDir, tt.args.ext, tt.args.devel)
			if tt.wantErr {
				require.Error(t, err)
			}

			require.NoError(t, err)

			var names []string
			for _, tmpl := range gotTmpl.Templates() {
				t.Log(tmpl.Name())
				names = append(names, tmpl.Name())
			}

			require.ElementsMatch(t, tt.templateNames, names )
			require.Equal(t, len(tt.templateNames), len(names), )


			//if !reflect.DeepEqual(gotTmpl, tt.wantTmpl) {
			//	t.Errorf("NewTemplateManager() gotTmpl = %v, want %v", gotTmpl, tt.wantTmpl)
			//}
		})
	}
}

func TestRender(t *testing.T) {
	type args struct {
		dir   string
		ext   string
		devel bool
		layoutDir string
	}

	type viewData struct {
		Title string
		ID string
	}

	tests := []struct {
		name     string
		args     args
		wantErr  bool
		template string
		data     interface{}
		content  []string
	}{
		{
			name:          "Render Index",
			args:          args{dir: "../../templates", layoutDir: "../../templates/layouts", ext: ".gohtml", devel: true},
			wantErr:       false,
			template: "index",
			data: nil,
			content: []string{
				"This is the Admin Index page",
				"My default sidebar content",
				"Index - Admin",
				"This is the Admin Footer",
			},
		},
		{
			name:          "Render Accounts/Index",
			args:          args{dir: "../../templates", layoutDir: "../../templates/layouts", ext: ".gohtml", devel: true},
			wantErr:       false,
			template: "accounts/index",
			data: &viewData{Title: "Accounts"},
			content: []string{
				"This is the accounts index page",
				"My default sidebar content",
				"Accounts - Admin",
				"Look I have a custom footer",
			},
		},
		{
			name:          "Render Users/Show",
			args:          args{dir: "../../templates", layoutDir: "../../templates/layouts", ext: ".gohtml", devel: true},
			wantErr:       false,
			template: "users/show",
			data: &viewData{Title: "Users 123", ID: "123"},
			content: []string{
				"Users Show 123",
				"I am in the user's show page",
				"My default sidebar content",
				"Show User 123",
				"This is the Admin Footer",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTmpl, err := NewTemplateManager(tt.args.dir, tt.args.layoutDir, tt.args.ext, tt.args.devel)
			if tt.wantErr {
				require.Error(t, err)
			}

			require.NoError(t, err)
			buf := new(bytes.Buffer)
			err = gotTmpl.Render(buf, tt.template, tt.data )
			require.NoError(t, err)
			output := buf.String()
			for _, content := range tt.content {
				require.Contains(t, output,content)
			}
			
		})
	}

}