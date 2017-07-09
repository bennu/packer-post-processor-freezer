package main

import (
	"context"
	"fmt"

	"github.com/hashicorp/packer/packer"
	"github.com/mitchellh/multistep"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
)

type StepMoveTemplate struct {
	Folder string
}

func (s *StepMoveTemplate) Run(state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)
	ctx := state.Get("context").(context.Context)
	finder := state.Get("finder").(*find.Finder)
	vm := state.Get("vm").(*object.VirtualMachine)
	d := state.Get("datacenter").(string)

	if s.Folder != "" {
		ui.Say("Moving template...")

		folder, err := finder.Folder(ctx, fmt.Sprintf("/%v/vm/%v", d, s.Folder))
		if err != nil {
			state.Put("error", err)
			ui.Error(err.Error())

			return multistep.ActionHalt
		}

		task, err := folder.MoveInto(ctx, []types.ManagedObjectReference{vm.Reference()})
		if err != nil {
			state.Put("error", err)
			ui.Error(err.Error())
			return multistep.ActionHalt
		}
		if err = task.Wait(ctx); err != nil {
			state.Put("error", err)
			ui.Error(err.Error())
			return multistep.ActionHalt
		}
	}
	return multistep.ActionContinue
}

func (s *StepMoveTemplate) Cleanup(multistep.StateBag) {}
