permit (
    principal is CareGiver,
    action,
    resource == Capability::"UserProfile"
)
when {
    context.impersonation == true &&
    principal.cr.contains(context.crId)&&
    context.CRCGActionSet.contains(action.name)
    // &&
    // context.CRCGResourceSet.contains(resource.name)
};
