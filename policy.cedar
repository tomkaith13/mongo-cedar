permit (
    principal is CareGiver,
    action,
    resource
)
when {
    context.impersonation == true &&
    principal.status == "ACCEPTED" &&
    principal.cr.contains(context.crId) &&
    context.CRCGActionSet.contains(action.name) &&
    context.CRCGResourceSet.contains(resource.name)
};
