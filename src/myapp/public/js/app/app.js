

var BookTab = new bookTab();
var StaffTab = new staffTab;
var JournalTab = new journalTab;


function initch(a) {
  return a;
}

function init() {
  var x = BookTab.getView();
  var y = StaffTab.getView();
  var z = JournalTab.getView();
  
 viewer = {
 rows:[
   {cols:[
    {view: "button", autowidth:true, align:"center", css:"webix_primary", id:"exit", value:"Выход", click:function(id,event){
      webix.ajax().get("/Exite").then(function(){
      window.location.reload();
      });
    }},
    {view: "button", autowidth:true, align:"center", css:"webix_primary", id:"AddUser", value:"Добавить нового пользователя", click:function(id,event){
      $$("winAddUser").show();
    }},
    {view: "button", autowidth:true, align:"center", css:"webix_primary", id:"UpdateUser", value:"Редактировать пользователя", click:function(id,event){
      $$("winUpdateUser").show();
    }},
    {view: "label", width:150, label: "", align:"center" },
    {view: "label",  label: "Библиотека RainbowSoft", align:"center" }
   ]
 },
 
  { view:"tabview", //главная
    cells:[
      { id: "booksView", header:"Книги", body: x,},
      { id: "staffView", header:"Сотрудники", body: y,},
      { id: "journal", header:"Журнал", body: z }
    ],
    multiview:{animate:true}
  },
 ]}
return viewer
}

function run(){
  webix.ui(init()); // инициализация структуры страницы
  
  
  webix.ui(BookTab.initWindow()); // добавления модального окна
  webix.ui(StaffTab.initWindow());
  
  BookTab.editeEvents(BookTab); // прикрепление событий
  StaffTab.editeEvents(StaffTab);
  JournalTab.editeEvents(JournalTab);
  
  
};
run();

// Окно добавления пользователя

webix.ui({
  view:"window",
  id:"winAddUser",
  height:600,
  position:"center",
  zIndex:1000,
  modal:true, 
  head:{
    view:"toolbar", cols:[
      
      { view:"label", label: "Добавить нового пользователя", align:"center", },
      { view:"icon", id:"exitWinAddUser", icon:"mdi mdi-close", align: 'right', click:function(id,event){
        $$("winAddUser").hide();
      }}, 
    ]
  },

  body:{
    rows: [
        {view:"form", id:"formAddUser",      
        elements: [
          { view:"text", name:"Username", label:"Логин", invalidMessage: "Не валидный логин", bottomLabel:"Символы A-z и цифры 0-9"},
          { view:"text", type:"password", name:"Password", label:"Пароль", invalidMessage: "Не валидный пароль", bottomLabel:"Символы A-z и цифры 0-9"},
          { view:"button", id:"butAddUser", type:"icon", icon:"mdi mdi-check", click:function(id,event){
            if (!$$("formAddUser").validate()){
              return
            }
            item_data = $$("formAddUser").getValues();
            webix.ajax().headers({
              "Content-type":"application/json"
          }).post("/Add", JSON.stringify(item_data)).then(function(data){
            data = data.text();
            console.log(data)
            if(data == "2") {
              webix.message({ type:"error", text:"Данный логин уже занят" });
              $$("winAddUser").show();
              return false;
            }
          });
            $$("winAddUser").hide();
            $$("formAddUser").clear();
          }}
          ],
          rules: {
            "Username":webix.rules.isNotEmpty,
            "Username":function(value){
              if (value.length < 5) return false;
              for (let char of value) {
                if (char<"0" || ((char > "9") && (char < "A")) || char > "z") {
                  return false
                }
              }
              return true
            },
            "Password":webix.rules.isNotEmpty,
            "Password":function(value){
              if (value.length < 5) return false;
              for (let char of value) {
                if (char<"0" || ((char > "9") && (char < "A")) || char > "z") {
                  return false
                }
              }
              return true
            },
          }
        },

    ]
     
  } 
});

// Окно обновления пользователя

webix.ui({
  view:"window",
  id:"winUpdateUser",
  height:600,
  position:"center",
  zIndex:1000,
  modal:true, 
  head:{
    view:"toolbar", cols:[
      
      { view:"label", label: "Редактировать пользователя", align:"center", },
      { view:"icon", id:"exitUpdateUser", icon:"mdi mdi-close", align: 'right', click:function(id,event){
        $$("winUpdateUser").hide();
      }}, 
    ]
  },

  body:{
    rows: [
        {view:"form", id:"formUpdateUser",      
        elements: [
          { view:"text", name:"Username", label:"Логин", invalidMessage: "Не валидный логин", bottomLabel:"Символы A-z и цифры 0-9"},
          { view:"text", type:"password", name:"Password", label:"Пароль", invalidMessage: "Не валидный пароль", bottomLabel:"Символы A-z и цифры 0-9" },
          { view:"button", id:"butUpdateUser", type:"icon", icon:"mdi mdi-check", click:function(id,event){
            if (!$$("formUpdateUser").validate()){
              return
            }
            item_data = $$("formUpdateUser").getValues();
            webix.ajax().headers({
              "Content-type":"application/json"
          }).post("/Update", JSON.stringify(item_data)).then(function(data){
            data = data.text();
            console.log(data)
            if(data == "2") {
              webix.message({ type:"error", text:"Данный логин уже занят" });
              $$("winUpdateUser").show();
              
              return false;
            }
          });
            $$("winUpdateUser").hide();
            $$("formUpdateUser").clear();
          }}
          ],
          rules: {
            "Username":webix.rules.isNotEmpty,
            "Username":function(value){
              if (value.length < 5) return false;
              for (let char of value) {
                if (char<"0" || ((char > "9") && (char < "A")) || char > "z") {
                  return false
                }
              }
              return true
            },
            "Password":webix.rules.isNotEmpty,
            "Password":function(value){
              if (value.length < 5) return false;
              for (let char of value) {
                if (char<"0" || ((char > "9") && (char < "A")) || char > "z") {
                  return false
                }
              }
              return true
            },
          }
        },

    ]
     
  } 
});







