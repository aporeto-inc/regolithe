package spec

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpecification_NewSpecification(t *testing.T) {

	Convey("Given I create a new specification", t, func() {

		spec := NewSpecification()

		Convey("Then the spec should be correctly initialized", func() {
			So(spec, ShouldNotBeNil)
		})
	})
}

func TestSpecification_Validate(t *testing.T) {

	Convey("Given I have a specification with no validation error", t, func() {

		s := &specification{
			RawModel: &Model{
				ResourceName: "things",
				RestName:     "thing",
				Description:  "desc.",
				EntityName:   "toto",
				Package:      "package",
			},
			RawAttributes: map[string][]*Attribute{
				"v1": []*Attribute{
					&Attribute{
						Name:        "attr1",
						Description: "desc.",
						Type:        "string",
					},
				},
			},
		}

		Convey("When I call validate", func() {

			err := s.Validate()

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a specification with validation error", t, func() {

		s := &specification{
			RawModel: &Model{
				ResourceName: "things",
				RestName:     "thing",
				Description:  "desc.",
				EntityName:   "toto",
			},
			RawAttributes: map[string][]*Attribute{
				"v1": []*Attribute{
					&Attribute{
						Identifier: false,
						Name:       "not-id",
					},
					&Attribute{
						Name:        "id",
						Type:        "coucou",
						Description: "wala",
					},
				},
			},
		}

		Convey("When I call validate", func() {

			err := s.Validate()

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, `thing.spec: schema error: attributes.v1.1.type: attributes.v1.1.type must be one of the following: "string", "integer", "float", "boolean", "enum", "list", "object", "time", "external"
thing.spec: schema error: description: description is required
thing.spec: schema error: package: package is required
thing.spec: schema error: type: type is required
thing.spec: schema error: v1: Additional property v1 is not allowed`)
			})
		})
	})

	Convey("Given I have an abstract with no validation error", t, func() {

		s := &specification{
			RawAttributes: map[string][]*Attribute{
				"v1": []*Attribute{
					&Attribute{
						Name:        "attr1",
						Description: "desc.",
						Type:        "string",
					},
				},
			},
		}

		Convey("When I call validate", func() {

			err := s.Validate()

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given I have an abstract with validation error", t, func() {

		s := &specification{
			RawAttributes: map[string][]*Attribute{
				"v1": []*Attribute{
					&Attribute{
						Name:        "attr1",
						Description: "desc.",
						Type:        "string",
					},
				},
			},
			RawRelations: []*Relation{
				&Relation{
					RestName: "a",
				},
			},
		}

		Convey("When I call validate", func() {

			err := s.Validate()

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, ".: schema error: relations: Additional property relations is not allowed")
			})
		})
	})
}

func TestSpecification_Getters(t *testing.T) {

	Convey("Given I have a new API", t, func() {

		s := &specification{
			RawModel: &Model{
				EntityName:   "Test",
				ResourceName: "tests",
				RestName:     "test",
				AllowsCreate: true,
				AllowsDelete: true,
				AllowsGet:    true,
				AllowsUpdate: true,
			},
			RawAttributes: map[string][]*Attribute{
				"v1": []*Attribute{
					&Attribute{
						Identifier: false,
						Name:       "not-id",
					},
					&Attribute{
						Identifier: true,
						Name:       "id",
					},
				},
			},
		}

		s.buildAttributesInfo() // nolint: errcheck

		Convey("Then the getters should work", func() {
			So(s.Identifier().Name, ShouldEqual, "id")
		})
	})
}

func TestSpecification_TypeProviders(t *testing.T) {

	Convey("Given I have a new API", t, func() {

		s := &specification{
			RawAttributes: map[string][]*Attribute{
				"v1": []*Attribute{
					&Attribute{
						Name:          "not-id",
						ConvertedName: "not-id",
						TypeProvider:  "toto",
					},
					&Attribute{
						ConvertedName: "id",
						TypeProvider:  "titi",
					},
					&Attribute{
						ConvertedName: "id2",
						TypeProvider:  "titi",
					},
					&Attribute{},
				},
			},
		}

		s.buildAttributesInfo() // nolint: errcheck

		Convey("When I call TypeProviders", func() {

			providers := s.TypeProviders()

			Convey("Then the providers should be correct", func() {
				So(providers, ShouldResemble, []string{"toto", "titi"})
			})
		})
	})
}

func TestSpecification_AttributeInitializers(t *testing.T) {

	Convey("Given I have a new API", t, func() {

		s := &specification{
			RawAttributes: map[string][]*Attribute{
				"v1": []*Attribute{
					&Attribute{
						Identifier:    false,
						Name:          "not-id",
						ConvertedName: "not-id",
						DefaultValue:  "default1",
						Type:          AttributeTypeString,
					},
					&Attribute{
						Identifier:    true,
						Name:          "id",
						ConvertedName: "id",
						Initializer:   "init",
					},
				},
			},
		}

		s.buildAttributesInfo() // nolint: errcheck

		Convey("When I call AttributeInitializers", func() {

			inits := s.AttributeInitializers("v1")

			Convey("Then the initializers should be correct", func() {
				So(inits["id"], ShouldEqual, "init")
				So(inits["not-id"], ShouldEqual, `"default1"`)
				So(len(inits), ShouldEqual, 2)
			})
		})
	})
}

func TestSpecification_OrderingAttributes(t *testing.T) {

	Convey("Given I have a new API", t, func() {

		s := &specification{
			RawAttributes: map[string][]*Attribute{
				"v1": []*Attribute{

					&Attribute{
						DefaultOrder: true,
						Name:         "a1",
					},
					&Attribute{
						DefaultOrder: true,
						Name:         "a2",
					},
					&Attribute{
						Name: "a3",
					},
				},
			},
		}

		s.buildAttributesInfo() // nolint: errcheck

		Convey("When I call OrderingAttributes", func() {

			o := s.OrderingAttributes("v1")

			Convey("Then the orderingAttributes should be correct", func() {
				So(len(o), ShouldEqual, 2)
				So(o[0].Name, ShouldEqual, "a1")
				So(o[1].Name, ShouldEqual, "a2")
			})
		})
	})
}

func TestSpecification_AttributeMap(t *testing.T) {

	Convey("Given I load the task specification file", t, func() {

		spec, err := LoadSpecification("./tests/task.spec", true)

		Convey("Then err should be nil", func() {
			So(err, ShouldBeNil)
		})

		Convey("Then the attribute map should be correctly built", func() {
			So(len(spec.Attributes("v1")), ShouldEqual, 3)
			So(spec.Attribute("name", "v1").Name, ShouldEqual, "name")
			So(spec.Attribute("description", "v1").Name, ShouldEqual, "description")
			So(spec.Attribute("status", "v1").Name, ShouldEqual, "status")
		})
	})
}

func TestSpecification_BuildAttributeNames(t *testing.T) {

	Convey("Given I create a specification with the same attribute twice.", t, func() {

		spec := &specification{
			RawAttributes: map[string][]*Attribute{
				"v1": []*Attribute{

					&Attribute{
						Name: "a",
					},
					&Attribute{
						Name: "a",
					},
				},
			},
		}

		Convey("When I call BuildAttributeNames", func() {

			err := spec.buildAttributesInfo()

			Convey("Then err Should Not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestSpecification_APIMap(t *testing.T) {

	Convey("Given I load the root specification file", t, func() {

		spec, err := LoadSpecification("./tests/root.spec", true)

		Convey("Then err should be nil", func() {
			So(err, ShouldBeNil)
		})

		Convey("Then the relation map should be correctly built", func() {
			So(len(spec.Relations()), ShouldEqual, 2)
			So(spec.Relation("list").RestName, ShouldEqual, "list")
			So(spec.Relation("user").RestName, ShouldEqual, "user")
		})
	})
}

func TestSpecification_buildRelationssInfo(t *testing.T) {

	Convey("Given I create a specification with the same relation twice.", t, func() {

		spec := &specification{
			RawRelations: []*Relation{
				&Relation{
					RestName: "a",
				},
				&Relation{
					RestName: "a",
				},
			},
		}

		Convey("When I call buildRelationssInfo", func() {

			err := spec.buildRelationssInfo()

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestSpecification_LoadSpecification(t *testing.T) {

	Convey("Given I load a non existing specification file", t, func() {

		_, err := LoadSpecification("./tests/not.spec", true)

		Convey("Then err should not be nil", func() {
			So(err, ShouldNotBeNil)
		})
	})

	Convey("Given I load a bad formatted specification file", t, func() {

		_, err := LoadSpecification("./tests/task.spec.bad", true)

		Convey("Then err should not be nil", func() {
			So(err, ShouldNotBeNil)
		})
	})

	Convey("Given I load the root specification file", t, func() {

		spec, err := LoadSpecification("./tests/root.spec", true)
		rels := spec.Relations()

		Convey("Then err should be nil", func() {
			So(err, ShouldBeNil)
		})

		Convey("Then the spec should be correctly initialized", func() {

			model := spec.Model()

			So(model.AllowsGet, ShouldBeTrue)
			So(model.AllowsCreate, ShouldBeFalse)
			So(model.AllowsDelete, ShouldBeFalse)
			So(model.AllowsUpdate, ShouldBeFalse)
			So(model.Description, ShouldEqual, "Root object of the API.")
			So(model.EntityName, ShouldEqual, "Root")
			So(model.Package, ShouldEqual, "todo-list")
			So(model.ResourceName, ShouldEqual, "root")
			So(model.RestName, ShouldEqual, "root")
			So(model.Extends, ShouldBeNil)
			So(model.IsRoot, ShouldBeTrue)
			So(model.Aliases, ShouldBeNil)
		})

		Convey("Then the number of relations should be correct", func() {
			So(len(rels), ShouldEqual, 2)
		})

		Convey("Then the list of relations should be correct", func() {
			So(rels[0].AllowsCreate, ShouldBeTrue)
			So(rels[0].AllowsDelete, ShouldBeFalse)
			So(rels[0].Deprecated, ShouldBeFalse)
			So(rels[0].AllowsGet, ShouldBeTrue)
			So(rels[0].RestName, ShouldEqual, "list")
			So(rels[0].AllowsUpdate, ShouldBeFalse)
		})

		Convey("Then the user relation should be correct", func() {
			So(rels[1].AllowsCreate, ShouldBeTrue)
			So(rels[1].AllowsDelete, ShouldBeFalse)
			So(rels[1].Deprecated, ShouldBeFalse)
			So(rels[1].AllowsGet, ShouldBeTrue)
			So(rels[1].RestName, ShouldEqual, "user")
			So(rels[1].AllowsUpdate, ShouldBeFalse)
		})

	})

	Convey("Given I load the task specification file", t, func() {

		spec, err := LoadSpecification("./tests/task.spec", true)
		attrs := spec.Attributes("v1")
		model := spec.Model()

		Convey("Then err should be nil", func() {
			So(err, ShouldBeNil)
		})

		Convey("Then the spec should be correctly initialized", func() {
			So(model.AllowsGet, ShouldBeTrue)
			So(model.AllowsCreate, ShouldBeFalse)
			So(model.AllowsDelete, ShouldBeTrue)
			So(model.AllowsUpdate, ShouldBeTrue)
			So(model.Description, ShouldEqual, "Represent a task to do in a listd.")
			So(model.EntityName, ShouldEqual, "Task")
			So(model.Package, ShouldEqual, "todo-list")
			So(model.ResourceName, ShouldEqual, "tasks")
			So(model.RestName, ShouldEqual, "task")
			So(model.Extends, ShouldResemble, []string{"@base"})
			So(model.IsRoot, ShouldBeFalse)
			So(model.Aliases, ShouldResemble, []string{"tsk"})
		})

		Convey("Then the number of attributes should be correct", func() {
			So(len(spec.Attributes("v1")), ShouldEqual, 3)
		})

		Convey("Then the spec attribute description be correctly initialized", func() {

			So(attrs[0].AllowedChars, ShouldBeEmpty)
			So(attrs[0].AllowedChoices, ShouldBeEmpty)
			So(attrs[0].Autogenerated, ShouldBeFalse)
			So(attrs[0].CreationOnly, ShouldBeFalse)
			So(attrs[0].DefaultOrder, ShouldBeFalse)
			So(attrs[0].DefaultValue, ShouldBeNil)
			So(attrs[0].Deprecated, ShouldBeFalse)
			So(attrs[0].Description, ShouldEqual, "The description.")
			So(attrs[0].Exposed, ShouldBeTrue)
			So(attrs[0].Filterable, ShouldBeTrue)
			So(attrs[0].ForeignKey, ShouldBeFalse)
			So(attrs[0].Format, ShouldEqual, AttributeFormatFree)
			So(attrs[0].Getter, ShouldBeFalse)
			So(attrs[0].Identifier, ShouldBeFalse)
			So(attrs[0].Index, ShouldBeFalse)
			So(attrs[0].MaxLength, ShouldEqual, 0)
			So(attrs[0].MaxValue, ShouldEqual, 0.0)
			So(attrs[0].MinLength, ShouldEqual, 0)
			So(attrs[0].MinValue, ShouldEqual, 0.0)
			So(attrs[0].Name, ShouldEqual, "description")
			So(attrs[0].Orderable, ShouldBeTrue)
			So(attrs[0].PrimaryKey, ShouldBeFalse)
			So(attrs[0].ReadOnly, ShouldBeFalse)
			So(attrs[0].Required, ShouldBeFalse)
			So(attrs[0].Secret, ShouldBeFalse)
			So(attrs[0].Setter, ShouldBeFalse)
			So(attrs[0].Stored, ShouldBeTrue)
			So(attrs[0].SubType, ShouldEqual, "")
			So(attrs[0].Transient, ShouldBeFalse)
			So(attrs[0].Type, ShouldEqual, AttributeTypeString)
		})

		Convey("Then the spec attribute name be correctly initialized", func() {

			So(attrs[1].AllowedChars, ShouldBeEmpty)
			So(attrs[1].AllowedChoices, ShouldBeEmpty)
			So(attrs[1].Autogenerated, ShouldBeFalse)
			So(attrs[1].CreationOnly, ShouldBeFalse)
			So(attrs[1].DefaultOrder, ShouldBeFalse)
			So(attrs[1].DefaultValue, ShouldBeNil)
			So(attrs[1].Deprecated, ShouldBeFalse)
			So(attrs[1].Description, ShouldEqual, "The name.")
			So(attrs[1].Exposed, ShouldBeTrue)
			So(attrs[1].Filterable, ShouldBeTrue)
			So(attrs[1].ForeignKey, ShouldBeFalse)
			So(attrs[1].Format, ShouldEqual, AttributeFormatFree)
			So(attrs[1].Getter, ShouldBeTrue)
			So(attrs[1].Identifier, ShouldBeFalse)
			So(attrs[1].Index, ShouldBeFalse)
			So(attrs[1].MaxLength, ShouldEqual, 0)
			So(attrs[1].MaxValue, ShouldEqual, 0.0)
			So(attrs[1].MinLength, ShouldEqual, 0)
			So(attrs[1].MinValue, ShouldEqual, 0.0)
			So(attrs[1].Name, ShouldEqual, "name")
			So(attrs[1].Orderable, ShouldBeTrue)
			So(attrs[1].PrimaryKey, ShouldBeFalse)
			So(attrs[1].ReadOnly, ShouldBeFalse)
			So(attrs[1].Required, ShouldBeTrue)
			So(attrs[1].Secret, ShouldBeFalse)
			So(attrs[1].Setter, ShouldBeTrue)
			So(attrs[1].Stored, ShouldBeTrue)
			So(attrs[1].SubType, ShouldEqual, "")
			So(attrs[1].Transient, ShouldBeFalse)
			So(attrs[1].Type, ShouldEqual, AttributeTypeString)
		})

		Convey("Then the spec attribbute status be correctly initialized", func() {

			So(attrs[2].AllowedChars, ShouldBeEmpty)
			So(attrs[2].AllowedChoices, ShouldResemble, []string{"DONE", "PROGRESS", "TODO"})
			So(attrs[2].Autogenerated, ShouldBeFalse)
			So(attrs[2].CreationOnly, ShouldBeFalse)
			So(attrs[2].DefaultOrder, ShouldBeFalse)
			So(attrs[2].DefaultValue, ShouldResemble, "TODO")
			So(attrs[2].Deprecated, ShouldBeFalse)
			So(attrs[2].Description, ShouldEqual, "The status of the task.")
			So(attrs[2].Exposed, ShouldBeTrue)
			So(attrs[2].Filterable, ShouldBeTrue)
			So(attrs[2].ForeignKey, ShouldBeFalse)
			So(attrs[2].Format, ShouldEqual, "")
			So(attrs[2].Getter, ShouldBeFalse)
			So(attrs[2].Identifier, ShouldBeFalse)
			So(attrs[2].Index, ShouldBeFalse)
			So(attrs[2].MaxLength, ShouldEqual, 0)
			So(attrs[2].MaxValue, ShouldEqual, 0.0)
			So(attrs[2].MinLength, ShouldEqual, 0)
			So(attrs[2].MinValue, ShouldEqual, 0.0)
			So(attrs[2].Name, ShouldEqual, "status")
			So(attrs[2].Orderable, ShouldBeTrue)
			So(attrs[2].PrimaryKey, ShouldBeFalse)
			So(attrs[2].ReadOnly, ShouldBeFalse)
			So(attrs[2].Required, ShouldBeFalse)
			So(attrs[2].Secret, ShouldBeFalse)
			So(attrs[2].Setter, ShouldBeFalse)
			So(attrs[2].Stored, ShouldBeTrue)
			So(attrs[2].SubType, ShouldEqual, "")
			So(attrs[2].Transient, ShouldBeFalse)
			So(attrs[2].Type, ShouldEqual, AttributeTypeEnum)
		})

		Convey("When I apply the base specification", func() {

			base, err := LoadSpecification("./tests/@base.abs", true)
			spec.ApplyBaseSpecifications(base) // nolint: errcheck

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the number of attributes should be correct", func() {
				So(len(spec.Attributes("v1")), ShouldEqual, 6)
			})

			Convey("Then the additional attributes should have been applied", func() {
				So(spec.Attribute("ID", "v1").Name, ShouldEqual, "ID")
				So(spec.Attribute("ID", "v1").Autogenerated, ShouldBeTrue)
				So(spec.Attribute("ID", "v1").Description, ShouldEqual, "The identifier.")
				So(spec.Attribute("ID", "v1").Identifier, ShouldBeTrue)
				So(spec.Attribute("ID", "v1").PrimaryKey, ShouldBeTrue)
				So(spec.Attribute("ID", "v1").ReadOnly, ShouldBeTrue)

				So(spec.Attribute("parentID", "v1").Name, ShouldEqual, "parentID")
				So(spec.Attribute("parentID", "v1").Autogenerated, ShouldBeTrue)
				So(spec.Attribute("parentID", "v1").Description, ShouldEqual, "The identifier of the parent of the object.")
				So(spec.Attribute("parentID", "v1").ForeignKey, ShouldBeTrue)
				So(spec.Attribute("parentID", "v1").Identifier, ShouldBeFalse)
				So(spec.Attribute("parentID", "v1").PrimaryKey, ShouldBeFalse)
				So(spec.Attribute("parentID", "v1").ReadOnly, ShouldBeTrue)
			})
		})
	})
}
