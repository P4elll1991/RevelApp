

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
   {view: "button", css:"webix_primary", id:"exit", inputWidth:100,  value:"Выход", click:function(id,event){
    webix.ajax().get("/Exite").then(function(){
    window.location.reload();
    });
}},
  { view:"tabview",
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
  webix.ui(init());
  
  
  webix.ui(BookTab.initWindow());
  webix.ui(StaffTab.initWindow());
  
  BookTab.editeEvents(BookTab);
  StaffTab.editeEvents(StaffTab);
  JournalTab.editeEvents(JournalTab);
  
  
};
run();





