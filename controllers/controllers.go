package controllers

import (
	"fmt"
	"strconv"
	"github.com/gin-gonic/gin"

  . "github.com/apdaza/oasRuler/models"
  . "github.com/apdaza/oasRuler/data"
  . "github.com/apdaza/oasRuler/utils"
)

var dbmap = InitDb()

/*GetDomains funcion para obtener todos los dominios*/
func GetDomains(c *gin.Context) {
  var domains []Domain
  _, err := dbmap.Select(&domains, "SELECT * FROM domain")
  if err == nil {
    c.JSON(200, domains)
  } else {
    c.JSON(404, gin.H{"error": "no domain(s) into the table"})
  }
  // curl -i http://localhost:8080/api/rules/domains
}

/*GetDomain funcion para obtener un lo dominio*/
func GetDomain(c *gin.Context) {
  id := c.Params.ByName("id")
  var domain Domain
  err := dbmap.SelectOne(&domain, "SELECT * FROM domain WHERE id=?", id)
  if err == nil {
    domain_id, _ := strconv.ParseInt(id, 0, 64)
    content := &Domain{
      Id:   domain_id,
      Name: domain.Name,
    }
    c.JSON(200, content)
  } else {
    c.JSON(404, gin.H{"error": "domain not found"})
  }
  // curl -i http://localhost:8080/api/rules/domains/1
}

/*PostDomain funcion para insertar un dominio*/
func PostDomain(c *gin.Context) {
  var domain Domain
  c.Bind(&domain)
  if domain.Name != "" {
    if insert, _ := dbmap.Exec(`INSERT INTO domain (name) VALUES (?)`, domain.Name); insert != nil {
      domain_id, err := insert.LastInsertId()
      if err == nil {
        content := &Domain{
          Id:   domain_id,
          Name: domain.Name,
        }
        c.JSON(201, content)
      } else {
        CheckErr(err, "Insert failed")
      }
    }
  } else {
    c.JSON(422, gin.H{"error": "fields are empty"})
  }
  // curl -i -X POST -H "Content-Type: application/json" -d "{ \"name\": \"Titan\" }" http://localhost:8080/api/rules/domains
}

/*UpdateDomain funcion para*/
func UpdateDomain(c *gin.Context) {
  id := c.Params.ByName("id")
  var domain Domain
  err := dbmap.SelectOne(&domain, "SELECT * FROM domain WHERE id=?", id)
  if err == nil {
    var json Domain
    c.Bind(&json)
    domain_id, _ := strconv.ParseInt(id, 0, 64)
    domain := Domain{
      Id:   domain_id,
      Name: json.Name,
    }
    if domain.Name != "" {
      _, err = dbmap.Update(&domain)
      if err == nil {
        c.JSON(200, domain)
      } else {
        CheckErr(err, "Updated failed")
      }
    } else {
      c.JSON(422, gin.H{"error": "fields are empty"})
    }
  } else {
    c.JSON(404, gin.H{"error": "domain not found"})
  }
  // curl -i -X PUT -H "Content-Type: application/json" -d "{ \"name\": \"Academica\" }" http://localhost:8080/api/rules/domains/1
}

/*DeleteDomain funcion para*/
func DeleteDomain(c *gin.Context) {
  id := c.Params.ByName("id")
  var domain Domain
  err := dbmap.SelectOne(&domain, "SELECT id FROM domain WHERE id=?", id)
  if err == nil {
    _, err = dbmap.Delete(&domain)
    if err == nil {
      c.JSON(200, gin.H{"id #" + id: " deleted"})
    } else {
      CheckErr(err, "Delete failed")
    }
  } else {
    c.JSON(404, gin.H{"error": "domain not found"})
  }
  // curl -i -X DELETE http://localhost:8080/api/rules/domains/1
}

/*GetComponents funcion para obtener componentes*/
func GetComponents(c *gin.Context) {
  var components []Component
  _, err := dbmap.Select(&components, "SELECT * FROM component")
  if err == nil {
    c.JSON(200, components)
  } else {
    c.JSON(404, gin.H{"error": "no component(s) into the table"})
  }
  // curl -i http://localhost:8080/api/rules/rules
}

/*GetComponentByRule funcion para obtener los componentes de una regla*/
func GetComponentByRule(c *gin.Context) {
  name := c.Params.ByName("name")
  var components []ComponentByRule
  _, err := dbmap.Select(&components, "select t.name as Comparator, c.path, c.value, c.literal "+
    " FROM rules.component as c "+
    " inner join rules.comparator as t on c.comparator = t.id "+
    " inner join rules.rule as r on r.id = c.rule "+
    " where r.name = ?", name)
  if err == nil {
    c.JSON(200, components)
  } else {
    fmt.Println(err)
    c.JSON(404, gin.H{"error": "no component(s) into the table by rule name: " + name})
  }
  // curl -i http://localhost:8080/api/rules/rules
}

/*GetComponent funcion para obtener un componente*/
func GetComponent(c *gin.Context) {
  id := c.Params.ByName("id")
  var component Component
  err := dbmap.SelectOne(&component, "SELECT * FROM component WHERE id=?", id)
  if err == nil {
    component_id, _ := strconv.ParseInt(id, 0, 64)
    content := &Component{
      Id:         component_id,
      Rule:       component.Rule,
      Comparator: component.Comparator,
      Path:       component.Path,
      Value:      component.Value,
      Literal:    component.Literal,
    }
    c.JSON(200, content)
  } else {
    c.JSON(404, gin.H{"error": "rule not found"})
  }
  // curl -i http://localhost:8080/api/rules/rules/1
}

/*PostComponent funcion para insertar*/
func PostComponent(c *gin.Context) {
  var component Component
  c.Bind(&component)
  if component.Rule > 0 && component.Comparator > 0 && component.Path != "" && component.Value != "" && component.Literal >= 0 {
    if insert, _ := dbmap.Exec(`INSERT INTO component (rule, comparator, path, value, literal)
                                VALUES (?, ?, ?, ?, ?, ?)`,
      component.Rule, component.Comparator, component.Path,
      component.Value, component.Literal); insert != nil {
      component_id, err := insert.LastInsertId()
      if err == nil {
        content := &Component{
          Id:         component_id,
          Rule:       component.Rule,
          Comparator: component.Comparator,
          Path:       component.Path,
          Value:      component.Value,
          Literal:    component.Literal,
        }
        c.JSON(201, content)
      } else {
        CheckErr(err, "Insert failed")
      }
    }
  } else {
    c.JSON(422, gin.H{"error": "fields are empty"})
  }
  // curl -i -X POST -H "Content-Type: application/json" -d "{ \"domain\": 1, \"name\": \"Tipo nomina\" }" http://localhost:8080/api/rules/rules
}

/*UpdateComponent funcion para*/
func UpdateComponent(c *gin.Context) {
  id := c.Params.ByName("id")
  var component Component
  err := dbmap.SelectOne(&component, "SELECT * FROM component WHERE id=?", id)
  if err == nil {
    var json Component
    c.Bind(&json)
    component_id, _ := strconv.ParseInt(id, 0, 64)
    component := Component{
      Id:         component_id,
      Rule:       json.Rule,
      Comparator: json.Comparator,
      Path:       json.Path,
      Value:      json.Value,
      Literal:    json.Literal,
    }
    if component.Rule > 0 && component.Comparator > 0 && component.Path != "" && component.Value != "" && component.Literal >= 0 {
      _, err = dbmap.Update(&component)
      if err == nil {
        c.JSON(200, component)
      } else {
        CheckErr(err, "Updated failed")
      }
    } else {
      c.JSON(422, gin.H{"error": "fields are empty"})
    }
  } else {
    c.JSON(404, gin.H{"error": "component not found"})
  }
  // curl -i -X PUT -H "Content-Type: application/json" -d "{ \"name\": \"Academica\" }" http://localhost:8080/api/rules/rules/1
}

/*DeleteComponent funcion para*/
func DeleteComponent(c *gin.Context) {
  id := c.Params.ByName("id")
  var component Component
  err := dbmap.SelectOne(&component, "SELECT id FROM component WHERE id=?", id)
  if err == nil {
    _, err = dbmap.Delete(&component)
    if err == nil {
      c.JSON(200, gin.H{"id #" + id: " deleted"})
    } else {
      CheckErr(err, "Delete failed")
    }
  } else {
    c.JSON(404, gin.H{"error": "component not found"})
  }
  // curl -i -X DELETE http://localhost:8080/api/rules/rules/1
}

/*GetRules funcion para obtener*/
func GetRules(c *gin.Context) {
  var rules []Rule
  _, err := dbmap.Select(&rules, "SELECT * FROM rule")
  if err == nil {
    c.JSON(200, rules)
  } else {
    c.JSON(404, gin.H{"error": "no rule(s) into the table"})
  }
  // curl -i http://localhost:8080/api/rules/rules
}

/*GetRule funcion para obtener un lo dominio*/
func GetRule(c *gin.Context) {
  id := c.Params.ByName("id")
  var rule Rule
  err := dbmap.SelectOne(&rule, "SELECT * FROM rule WHERE id=?", id)
  if err == nil {
    rule_id, _ := strconv.ParseInt(id, 0, 64)
    content := &Rule{
      Id:          rule_id,
      Domain:      rule.Domain,
      Name:        rule.Name,
      Description: rule.Description,
    }
    c.JSON(200, content)
  } else {
    c.JSON(404, gin.H{"error": "rule not found"})
  }
  // curl -i http://localhost:8080/api/rules/rules/1
}

/*PostRule funcion para insertar*/
func PostRule(c *gin.Context) {
  var rule Rule
  c.Bind(&rule)
  if rule.Name != "" && rule.Domain > 0 {
    if insert, _ := dbmap.Exec(`INSERT INTO rule (name, domain, description) VALUES (?, ?, ?)`, rule.Name, rule.Domain, rule.Description); insert != nil {
      rule_id, err := insert.LastInsertId()
      if err == nil {
        content := &Rule{
          Id:          rule_id,
          Domain:      rule.Domain,
          Name:        rule.Name,
          Description: rule.Description,
        }
        c.JSON(201, content)
      } else {
        CheckErr(err, "Insert failed")
      }
    }
  } else {
    c.JSON(422, gin.H{"error": "fields are empty"})
  }
  // curl -i -X POST -H "Content-Type: application/json" -d "{ \"domain\": 1, \"name\": \"Tipo nomina\" }" http://localhost:8080/api/rules/rules
}

/*UpdateRule funcion para*/
func UpdateRule(c *gin.Context) {
  id := c.Params.ByName("id")
  var rule Rule
  err := dbmap.SelectOne(&rule, "SELECT * FROM rule WHERE id=?", id)
  if err == nil {
    var json Rule
    c.Bind(&json)
    rule_id, _ := strconv.ParseInt(id, 0, 64)
    rule := Rule{
      Id:          rule_id,
      Domain:      json.Domain,
      Name:        json.Name,
      Description: json.Description,
    }
    if rule.Name != "" && rule.Domain > 0 {
      _, err = dbmap.Update(&rule)
      if err == nil {
        c.JSON(200, rule)
      } else {
        CheckErr(err, "Updated failed")
      }
    } else {
      c.JSON(422, gin.H{"error": "fields are empty"})
    }
  } else {
    c.JSON(404, gin.H{"error": "rule not found"})
  }
  // curl -i -X PUT -H "Content-Type: application/json" -d "{ \"name\": \"Academica\" }" http://localhost:8080/api/rules/rules/1
}

/*DeleteRule funcion para*/
func DeleteRule(c *gin.Context) {
  id := c.Params.ByName("id")
  var rule Rule
  err := dbmap.SelectOne(&rule, "SELECT id FROM rule WHERE id=?", id)
  if err == nil {
    _, err = dbmap.Delete(&rule)
    if err == nil {
      c.JSON(200, gin.H{"id #" + id: " deleted"})
    } else {
      CheckErr(err, "Delete failed")
    }
  } else {
    c.JSON(404, gin.H{"error": "rule not found"})
  }
  // curl -i -X DELETE http://localhost:8080/api/rules/rules/1
}
