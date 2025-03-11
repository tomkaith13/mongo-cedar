permit (
    principal is CareGiver,
    action,
    resource
)
when {
    context.impersonation == true &&
    principal.cr.contains(context.crId)&&
    context.CRCGActionSet.contains(action.name)
};
