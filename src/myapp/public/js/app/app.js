

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
  view:"tabview",
  cells:[
    { id: "booksView", header:"Книги", body: x,},
    { id: "staffView", header:"Сотрудники", body: y,},
    { id: "journal", header:"Журнал", body: z }
  ],
  multiview:{animate:true}
};
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





