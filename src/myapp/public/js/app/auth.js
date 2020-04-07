webix.ui({
    view:"window",
    id:"inter",
    height:600,
    position:"center",
    zIndex:1000,
    modal:true, 
    head:{
      view:"toolbar", cols:[
        
        { view:"label", label: "Вход", align:"center", },
      ]
    },

    body:{
      rows: [
          {view:"form", id:"formInter",      
          elements: [
            { view:"text", name:"username", label:"Логин",},
            { view:"text", type:"password", name:"pass", label:"Пароль" },
            { view:"button", id:"interBut", type:"icon",label:"Войти",  icon:"mdi mdi-check", }
            ],
          },

      ]
       
    } 
}).show();

$$("interBut").attachEvent("onItemClick", function(){
    var form = $$("formInter");
    var item_data = form.getValues();
    this.postData = {
        action:"info",
        Username:item_data.username, 
        Pass:item_data.pass, 
        }
        console.log(this.postData)
    webix.ajax().headers({
        "Content-type":"application/json"
    }).get("/Inter", JSON.stringify(this.postData)).then(function(){
        window.location.reload();
    });

});